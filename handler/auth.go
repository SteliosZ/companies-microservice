package handler

import (
	"company/microservice/model"
	"company/microservice/repository"
	"company/microservice/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type UserHandler struct {
	UserRepository repository.IUserRepository
}

func NewUserHandler(repo repository.IUserRepository) IUserHandler {
	return &UserHandler{UserRepository: repo}
}

// Login handler contains the functionality of the "Login" feature
func (h UserHandler) Login(c *fiber.Ctx) error {
	user := model.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong input data",
			"data":    err,
		})
	}

	foundUser, err := h.UserRepository.GetUserByEmail(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "user or password wrong",
			"data":    err,
		})
	}

	token, err := utils.GenerateJWT(&foundUser.ID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"status":  "error",
				"message": "user or password wrong",
				"data":    err,
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  "success",
			"message": "user registered",
			"data":    token,
		})
}

// Register handler contains the "Register" feature functionality
func (h UserHandler) Register(c *fiber.Ctx) error {

	newUser := model.User{}

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong input data",
			"data":    err,
		})
	}

	err := h.UserRepository.CreateUser(&newUser)
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
