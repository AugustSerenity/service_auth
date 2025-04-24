package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AugustSerenity/service_auth/internal/config"
)

func InitDB(cfg config.DB) *sql.DB {
	var err error
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Name)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Database connected!")
	return conn
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}
}
