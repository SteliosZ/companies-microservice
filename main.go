package main

import (
	"company/microservice/app"
	"company/microservice/config"
	"company/microservice/database"
	"log"
)

func main() {
	// load .env file
	err := config.Config()
	if err != nil {
		log.Fatalf("configuration not loaded: %v", err)
	}

	// // Initialize DB Connection
	tx := database.ConnectToDB()

	// Initialize Rest
	app.Setup(tx)
}
