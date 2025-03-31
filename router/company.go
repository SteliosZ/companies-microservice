package router

import (
	"company/microservice/handler"
	"company/microservice/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterCompanyRouter(api fiber.Router, h handler.ICompanyHandler) {
	// Company Routes
	company := api.Group("/company")
	company.Get("/:id", h.GetCompanyByID)

	company.Use(middleware.JWTCheck)
	company.Post("/", h.CreateCompany)
	company.Patch(":id", h.UpdateCompany)
	company.Delete("/:id", h.DeleteCompany)
}
