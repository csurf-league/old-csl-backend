package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SESSION_SECRET_KEY string
var SESSION_NAME string
var STEAM_API_KEY string
var PORT string

func GetEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	SESSION_SECRET_KEY = os.Getenv("SESSION_SECRET_KEY")
	SESSION_NAME = os.Getenv("SESSION_NAME")
	STEAM_API_KEY = os.Getenv("STEAM_API_KEY")
	PORT = os.Getenv("PORT")
}
