// Package config handles database configuration and initialization.
package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// NewGormDB initializes a new GORM database connection using the provided SQL connection.
func NewGormDB(conn *sql.DB) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	log.Println("Database connection established")
	return DB, nil
}

// getEnv retrieves the value of the environment variable key, or returns fallback if not set.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
