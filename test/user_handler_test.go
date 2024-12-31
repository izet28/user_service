// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gorilla/mux"
// 	"github.com/izet28/user_service/internal/handlers"
// 	"github.com/izet28/user_service/internal/models"
// 	"github.com/izet28/user_service/pkg/cache"
// 	"github.com/stretchr/testify/assert"
// )

// type MockService struct{}

// func (m *MockService) GetAllUsers() ([]models.User, error) {
// 	return []models.User{
// 		{ID: 1, Username: "JohnDoe", Email: "john@example.com"},
// 	}, nil
// }

// func (m *MockService) CreateUser(user *models.User) (*models.User, error) {
// 	user.ID = 1
// 	return user, nil
// }

// func (m *MockService) UpdateUser(id int, user *models.User) (*models.User, error) {
// 	user.ID = id
// 	return user, nil
// }

// func (m *MockService) DeleteUser(id int) error {
// 	return nil
// }

// func TestGetAllUsers(t *testing.T) {
// 	mockService := &MockService{}
// 	redisCache := cache.NewRedisCache("localhost:6379", "", 0)
// 	handler := handlers.NewUserHandler(mockService, redisCache)

// 	router := mux.NewRouter()
// 	router.HandleFunc("/users", handler.GetAllUsers).Methods("GET")

// 	req, _ := http.NewRequest("GET", "/users", nil)
// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// }

// func TestCreateUser(t *testing.T) {
// 	mockService := &MockService{}
// 	redisCache := cache.NewRedisCache("localhost:6379", "", 0)
// 	handler := handlers.NewUserHandler(mockService, redisCache)
// 	handler.Validate = validator.New()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/users", handler.CreateUser).Methods("POST")

// 	user := models.User{
// 		Username: "JohnDoe",
// 		Email:    "john@example.com",
// 		Password: "password123",
// 	}

// 	body, _ := json.Marshal(user)
// 	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusCreated, rr.Code)
// }
