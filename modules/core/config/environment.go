package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnvironment() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

}

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
