package config

import (
	"os"
)

func InitializeEnvironment() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }

}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
