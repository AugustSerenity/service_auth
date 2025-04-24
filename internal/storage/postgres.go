package storage

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB() *sql.DB {
	var err error
	connStr := "host=localhost port=5433 user=postgres dbname=medods sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	fmt.Println("Database connection Работает!!!")
	return conn
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}
}
