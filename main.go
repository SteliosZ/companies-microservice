package main

import (
	"company/microservice/database"
	"company/microservice/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Rest
	app := fiber.New(fiber.Config{
		AppName: "Company Handler Microservice",
	})

	// Initialize DB Connection
	database.ConnectToDB()

	// Setup Routes
	router.SetupRoutes(app)

	// Start Rest
	log.Fatal(app.Listen(":8080"))
}
