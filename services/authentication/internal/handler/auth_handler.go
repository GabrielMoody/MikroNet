package handler

import (
	"github.com/GabrielMoody/mikronet-auth-service/internal/controller"
	"github.com/GabrielMoody/mikronet-auth-service/internal/repository"
	"github.com/GabrielMoody/mikronet-auth-service/internal/service"
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
	authHandler.Post("/reset-password", authController.SendResetPasswordLink)
	authHandler.Put("/reset-password/:code", authController.ResetPassword)
	authHandler.Put("/change-password", authController.ChangePassword)
	authHandler.Get("/reset-password/:code", authController.ResetPasswordUI)
}
