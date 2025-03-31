package main

import (
	"company/microservice/config"
	"company/microservice/database"
	"company/microservice/router"
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

	// Initialize Handlers and Repositories

	// Initialize App
	router.InitializeApp(tx)
}
