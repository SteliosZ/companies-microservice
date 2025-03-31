package router

import (
	"company/microservice/database"
	"company/microservice/handler"
	"company/microservice/repository"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRouter(api fiber.Router, tx *database.DBInstance) {
	r := repository.NewUserRepository(tx.DB)
	h := handler.NewUserHandler(r)

	// JWT Middleware
	auth := api.Group("/auth")
	auth.Post("/login", h.Login)
	auth.Post("/register", h.Register)
}
