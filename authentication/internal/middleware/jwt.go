package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func errorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{
		Key: []byte(os.Getenv("JWT_SECRET")),
	},
		ErrorHandler: errorHandler,
	})
}
