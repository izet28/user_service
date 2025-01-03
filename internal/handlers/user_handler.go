package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/izet28/user_service/internal/models"
	"github.com/izet28/user_service/internal/repository"
	"github.com/izet28/user_service/internal/services"
	"github.com/izet28/user_service/pkg/cache"
	"github.com/izet28/user_service/pkg/logger"
	"github.com/izet28/user_service/pkg/utils"
)

type UserHandler struct {
	Service  *services.UserService
	Cache    *cache.RedisCache
	Validate *validator.Validate
}

func NewUserHandler(service *services.UserService, redisCache *cache.RedisCache) *UserHandler {
	return &UserHandler{
		Service:  service,
		Cache:    redisCache,
		Validate: validator.New(),
	}
}

// Get all users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	const cacheKey = "users"

	// Check cache
	cachedData, err := h.Cache.Get(cacheKey)
	if err == nil {
		utils.RespondWithJSON(w, http.StatusOK, cachedData)
		return
	}

	// Fetch from database
	users, err := h.Service.GetAllUsers()
	if err != nil {
		logger.Error("Failed to retrieve users: " + err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	// Cache data
	h.Cache.Set(cacheKey, users, 10*time.Minute)
	utils.RespondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Check cache
	userCacheKey := fmt.Sprintf("user:%d", id)
	cachedData, err := h.Cache.Get(userCacheKey)
	if err == nil {
		var user models.User
		if jsonErr := json.Unmarshal([]byte(cachedData), &user); jsonErr == nil {
			logger.Info("Data retrieved from Redis for key: " + userCacheKey)
			utils.RespondWithJSON(w, http.StatusOK, user)
			return
		}
	}

	// Fetch user from database
	logger.Info("Cache miss. Fetching user from database for ID: " + strconv.Itoa(id))
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		logger.Error("Failed to retrieve user from database: " + err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	// Cache the user data
	h.Cache.Set(userCacheKey, user, 10*time.Minute)
	logger.Info("User data cached in Redis for key: " + userCacheKey)
	utils.RespondWithJSON(w, http.StatusOK, user)
}

// Create a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Invalid input for CreateUser")
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Validate input
	if err := h.Validate.Struct(user); err != nil {
		logger.Error("Validation error in CreateUser: " + err.Error())
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create user
	createdUser, err := h.Service.CreateUser(&user)
	if err != nil {
		logger.Error("Failed to create user: " + err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	logger.Info("User created successfully: " + user.Username)
	h.Cache.Delete("users") // Clear cache
	utils.RespondWithJSON(w, http.StatusCreated, createdUser)
}

// Update a user
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Validate input
	if err := h.Validate.Struct(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Update user
	updatedUser, err := h.Service.UpdateUser(id, &user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	logger.Info("User updated successfully: " + user.Username)

	// Clear cache for this specific user
	userCacheKey := fmt.Sprintf("user:%d", id)
	h.Cache.Delete(userCacheKey)
	h.Cache.Delete("users") // Clear global cache

	utils.RespondWithJSON(w, http.StatusOK, updatedUser)
}

// Delete a user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.Service.DeleteUser(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	logger.Info("User deleted successfully: ID " + strconv.Itoa(id))

	// Clear cache for this specific user
	userCacheKey := fmt.Sprintf("user:%d", id)
	h.Cache.Delete(userCacheKey)
	h.Cache.Delete("users") // Clear global cache

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// Setup routes
func SetupRoutes(router *mux.Router, db *sql.DB, redisCache *cache.RedisCache) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService, redisCache)

	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	api.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	api.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
}
