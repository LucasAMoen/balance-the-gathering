package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key, fallback string) string {
	godotenv.Load()

	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
