// cmd/user-service/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/izet28/user_service/internal/model"
	"github.com/izet28/user_service/internal/transport"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := gorm.Open("postgres", "host=10.104.0.15 user=admin dbname=user password=admin_password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the database schema
	db.AutoMigrate(&model.User{})

	// Set up the router
	router := transport.NewRouter(db)

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
