package router

import (
	"company/microservice/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Health Check
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// TODO: JWT Middleware

	// Company Routes
	company := api.Group("/company")
	company.Get("/:name", handler.GetCompany)
	company.Post("/", handler.CreateCompany)
	company.Patch(":name", handler.UpdateCompany)
	company.Delete("/:name", handler.DeleteCompany)
}
