package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	env "github.com/robyzzz/csl-backend/config"
)

var db *sqlx.DB
var err error

func Connect() {
	db, err = sqlx.Open("postgres", connToString())
	if err != nil {
		log.Fatalf("Error connecting to the DB: %s\n", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error: Could not ping database: %s\n", err.Error())
	}

	db.MustExec(schema)
	log.Printf("Database connection done successfully: %p\n", db)

	users := []SteamUser{}
	err := db.Select(&users, "SELECT * FROM steam_user;")
	if err != nil {
		log.Fatalf("Error: Could not ping get player: %s\n", err.Error())
	}
	fmt.Println(users)
}

// Return a string for our db connection info
func connToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.POSTGRES_URL, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB)
}

// Returns date formatted as PQ's NOW()
func DBGetDateNow() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d:%04d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0)
}
