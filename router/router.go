package router

import (
	"company/microservice/handler"
	"company/microservice/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Health Check
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// JWT Middleware
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)

	// Company Routes
	company := api.Group("/company")
	company.Get("/:name", handler.GetCompanyByName)

	company.Use(middleware.JWTProtected)
	company.Post("/", handler.CreateCompany)
	company.Patch(":name", handler.UpdateCompany)
	company.Delete("/:name", handler.DeleteCompany)
}
