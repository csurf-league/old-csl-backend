package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SESSION_SECRET_KEY string
var SESSION_NAME string
var STEAM_API_KEY string
var PORT string

var POSTGRES_URL string
var POSTGRES_PORT string
var POSTGRES_USER string
var POSTGRES_PASSWORD string
var POSTGRES_DB string

func GetEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	SESSION_SECRET_KEY = os.Getenv("SESSION_SECRET_KEY")
	SESSION_NAME = os.Getenv("SESSION_NAME")
	STEAM_API_KEY = os.Getenv("STEAM_API_KEY")
	PORT = os.Getenv("PORT")

	POSTGRES_URL = os.Getenv("POSTGRES_URL")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
}
