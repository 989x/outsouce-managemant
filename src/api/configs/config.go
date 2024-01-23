package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from the .env file
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetPort retrieves the port from the environment variable or defaults to 8000
func GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	return port
}
