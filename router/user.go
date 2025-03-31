package router

import (
	"company/microservice/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRouter(api fiber.Router, h handler.IUserHandler) {
	auth := api.Group("/auth")
	auth.Post("/login", h.Login)
	auth.Post("/register", h.Register)
}
