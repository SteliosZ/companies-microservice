package router

import (
	"company/microservice/database"
	"company/microservice/handler"
	"company/microservice/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// RegisterRouters intializes every route
func RegisterRouters(app *fiber.App, tx *database.DBInstance) {
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

}
