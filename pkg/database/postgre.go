package database

import (
	"database/sql"
	"fmt"

	"github.com/izet28/user_service/config"

	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	return sql.Open("postgres", connStr)
}
