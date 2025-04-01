package app

import (
	"company/microservice/database"
	"company/microservice/router"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Setup initiates the fiber app
func Setup(tx *database.DBInstance) {
	// Initialize RestAPI
	app := fiber.New(fiber.Config{
		AppName: "Company Handler Microservice",
	})
	app.Use(cors.New())

	// Register Routes
	router.RegisterRouters(app, tx)

	// Handle Undefined Routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Start RestAPI
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("SERVICE_EXPOSED_PORT"))))
}
