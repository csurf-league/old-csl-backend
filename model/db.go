package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	env "github.com/robyzzz/csl-backend/config"
)

var db *sql.DB

func Connect() {
	db, err := sql.Open("postgres", connToString())
	if err != nil {
		log.Fatalf("Error connecting to the DB: %s\n", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error: Could not ping database: %s\n", err.Error())
	}

	log.Printf("Database connection done successfully\n")
}

// Return a string for our db connection info
func connToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.POSTGRES_URL, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB)
}
