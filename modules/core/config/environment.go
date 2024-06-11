package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnvironment() {
	godotenv.Load()

}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
