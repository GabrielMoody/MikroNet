package handler

import (
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthHandler(r fiber.Router, authController controller.AuthController) {
	authHandler := r.Group("/")

	authHandler.Post("/register/user", authController.CreateUser)
	authHandler.Post("/register/driver", authController.CreateDriver)
	authHandler.Post("/login", authController.LoginUser)

	authHandler.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{"status": "pass"})
	})
}
