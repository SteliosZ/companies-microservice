package main

import (
	"company/microservice/database"
	"company/microservice/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize Rest
	app := fiber.New(fiber.Config{
		AppName: "Company Handler Microservice",
	})
	app.Use(cors.New())

	// Initialize DB Connection
	database.Connect()

	// Setup Routes
	router.SetupRoutes(app)

	// Handle Undefined Routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Start Rest
	log.Fatal(app.Listen(":8080"))
}
