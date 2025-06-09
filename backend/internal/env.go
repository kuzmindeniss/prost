package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getIsProduction() bool {
	return os.Getenv("ENV") == "production"
}

func LoadEnv() {
	if getIsProduction() {
		return
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
