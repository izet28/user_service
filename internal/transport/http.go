// internal/transport/http.go
package transport

import (
	"github.com/gorilla/mux"
	"github.com/izet28/user_service/internal/handler"
	"github.com/izet28/user_service/internal/repository"
	"github.com/izet28/user_service/internal/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL driver
)

func NewRouter(db *gorm.DB) *mux.Router {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")

	return router
}
