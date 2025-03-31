package router

import (
	"company/microservice/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes intializes every route
func SetupRoutes(app *fiber.App, tx *database.DBInstance) {
	// Health Check
	api := app.Group("/api", logger.New())

	RegisterUserRouter(api, tx)
	RegisterCompanyRouter(api, tx)
}
