package middleware

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": "error",
		"error":  err.Error(),
	})
}

func JWTMiddleware() fiber.Handler {
	v := helper.LoadEnv()
	return jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{
		Key: []byte(v.GetString("JWT_SECRET")),
	},
		ErrorHandler: errorHandler,
	})
}
