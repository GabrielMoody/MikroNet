package handler

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/controller"
	"github.com/GabrielMoody/MikroNet/authentication/internal/repository"
	"github.com/GabrielMoody/MikroNet/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProfileHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewProfileRepo(db)
	profileService := service.NewProfileService(repo)
	profileController := controller.NewProfileController(profileService)

	profileHandler := r.Group("/auth")

	profileHandler.Post("/register", profileController.CreateUser)
	profileHandler.Post("/login", profileController.LoginUser)
	profileHandler.Post("/reset-password", profileController.SendResetPasswordLink)
	profileHandler.Put("/reset-password/:code", profileController.ResetPassword)
}
