package main

import (
	"company/microservice/config"
	"company/microservice/database"
	"company/microservice/router"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// load .env file
	err := config.Config()
	if err != nil {
		log.Fatalf("configuration not loaded: %v", err)
	}

	// Initialize Rest
	app := fiber.New(fiber.Config{
		AppName: "Company Handler Microservice",
	})
	app.Use(cors.New())

	// Initialize DB Connection
	database.ConnectToDB()

	// Setup Routes
	router.SetupRoutes(app)

	// Handle Undefined Routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Start Rest
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("SERVICE_EXPOSED_PORT"))))
}
