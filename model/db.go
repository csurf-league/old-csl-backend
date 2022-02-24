package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	env "github.com/robyzzz/csl-backend/config"
)

var db *sql.DB
var err error

func Connect() {
	db, err = sql.Open("postgres", connToString())
	if err != nil {
		log.Fatalf("Error connecting to the DB: %s\n", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error: Could not ping database: %s\n", err.Error())
	}

	_, err = db.Query(`CREATE TABLE IF NOT EXISTS steam_user (
		id SERIAL PRIMARY KEY,
		steamid TEXT,
		personaname TEXT,
		lastlogoff INTEGER,
		profileurl TEXT,
		avatar TEXT,
		avatarmedium TEXT,
		avatarfull TEXT,
		realname TEXT,
		primaryclanid TEXT,
		timecreated INTEGER,
		loccountrycode TEXT,
		gameid INTEGER  
	  );`)
	if err != nil {
		log.Fatalf("Error: Could not create table steam_users: %s\n", err.Error())
	}

	log.Printf("Database connection done successfully: %p\n", db)
}

// Return a string for our db connection info
func connToString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.POSTGRES_URL, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB)
}
