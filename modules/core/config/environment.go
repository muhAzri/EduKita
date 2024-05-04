package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
