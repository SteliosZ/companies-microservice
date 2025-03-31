package router

import (
	"company/microservice/database"
	"company/microservice/handler"
	"company/microservice/repository"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// InitializeApp intializes every route
func InitializeApp(tx *database.DBInstance) {
	// Initialize Rest
	app := fiber.New(fiber.Config{
		AppName: "Company Handler Microservice",
	})
	app.Use(cors.New())

	// User Handler Initialization
	userHandler := handler.NewUserHandler(
		repository.NewUserRepository(tx.DB),
	)

	// Company Handler Initialization
	companyHandler := handler.NewCompanyHandler(
		repository.NewCompanyRepository(tx.DB),
	)

	// Setup Routes
	api := app.Group("/api", logger.New())
	RegisterUserRouter(api, userHandler)
	RegisterCompanyRouter(api, companyHandler)

	// Handle Undefined Routes
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Start RestAPI
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("SERVICE_EXPOSED_PORT"))))
}
