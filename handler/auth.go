package handler

import (
	"company/microservice/database"
	"company/microservice/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return nil
}

func Register(c *fiber.Ctx) error {

	newUser := model.User{}

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong input data",
			"data":    err,
		})
	}

	err := database.CreateUser(newUser)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).
			JSON(fiber.Map{
				"status":  "error",
				"message": "user not created",
				"data":    err,
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  "success",
			"message": "user registered",
			"data":    err,
		})
}
