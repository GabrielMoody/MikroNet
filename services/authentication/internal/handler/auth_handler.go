package handler

import (
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/controller"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewAuthRepo(db)
	authService := service.NewAuthService(repo)
	authController := controller.NewAuthController(authService)

	authHandler := r.Group("/")

	authHandler.Post("/register/user", authController.CreateUser)
	authHandler.Post("/register/driver", authController.CreateDriver)
	authHandler.Post("/login", authController.LoginUser)

	authHandler.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{"status": "pass"})
	})
}
