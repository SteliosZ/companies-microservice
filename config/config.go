package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Config loads the .env configuration file
func Config() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("config load failure: %v", err.Error())
	}
	return err
}
