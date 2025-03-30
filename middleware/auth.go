package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func JWTCheck(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "authentication error",
				"data":    err.Error(),
			})
		},
	})(c)
}
