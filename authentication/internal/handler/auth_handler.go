package handler

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/controller"
	"github.com/GabrielMoody/MikroNet/authentication/internal/pb"
	"github.com/GabrielMoody/MikroNet/authentication/internal/repository"
	"github.com/GabrielMoody/MikroNet/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthHandler(r fiber.Router, db *gorm.DB, user pb.UserServiceClient, driver pb.DriverServiceClient, dashboard pb.OwnerServiceClient) {
	repo := repository.NewAuthRepo(db)
	authService := service.NewAuthService(repo)
	authController := controller.NewAuthController(authService, user, driver, dashboard)

	authHandler := r.Group("/")

	authHandler.Post("/register/user", authController.CreateUser)
	authHandler.Post("/register/driver", authController.CreateDriver)
	authHandler.Post("/register/owner", authController.CreateOwner)
	authHandler.Post("/login", authController.LoginUser)
	authHandler.Post("/reset-password", authController.SendResetPasswordLink)
	authHandler.Put("/reset-password/:code", authController.ResetPassword)
	authHandler.Put("/change-password", authController.ChangePassword)
}
