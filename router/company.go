package router

import (
	"company/microservice/database"
	"company/microservice/handler"
	"company/microservice/middleware"
	"company/microservice/repository"

	"github.com/gofiber/fiber/v2"
)

func RegisterCompanyRouter(api fiber.Router, tx *database.DBInstance) {
	r := repository.NewCompanyRepository(tx.DB)
	h := handler.NewCompanyHandler(r)

	// Company Routes
	company := api.Group("/company")
	company.Get("/:id", h.GetCompanyByID)

	company.Use(middleware.JWTCheck)
	company.Post("/", h.CreateCompany)
	company.Patch(":id", h.UpdateCompany)
	company.Delete("/:id", h.DeleteCompany)
}
