package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func loads .env configuration file
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("config load failure: %v", err.Error())
	}
	return os.Getenv(key)
}
