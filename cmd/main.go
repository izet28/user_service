package main

import (
	"net/http"

	"github.com/izet28/user_service/config"
	"github.com/izet28/user_service/internal/handlers"
	"github.com/izet28/user_service/pkg/cache"
	"github.com/izet28/user_service/pkg/database"
	"github.com/izet28/user_service/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi logger
	logger.InitLogger("app.log")

	logger.Info("Loading configuration...")
	cfg := config.LoadConfig()
	logger.Info("Configuration loaded successfully")

	logger.Info("Connecting to database...")
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("Could not connect to database: " + err.Error())
	}
	defer db.Close()
	logger.Info("Database connection successful")

	logger.Info("Connecting to Redis...")
	redisCache := cache.NewRedisCache(cfg.RedisAddr, cfg.RedisDB)
	logger.Info("Redis connection successful")

	logger.Info("Setting up routes...")
	router := mux.NewRouter()
	handlers.SetupRoutes(router, db, redisCache)

	logger.Info("Starting HTTP server on port " + cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, router)
	if err != nil {
		logger.Fatal("Failed to start HTTP server: " + err.Error())
	}
}
