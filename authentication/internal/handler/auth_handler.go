package handler

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/controller"
	"github.com/GabrielMoody/MikroNet/authentication/internal/pb"
	"github.com/GabrielMoody/MikroNet/authentication/internal/repository"
	"github.com/GabrielMoody/MikroNet/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProfileHandler(r fiber.Router, db *gorm.DB, client pb.UserServiceClient) {
	repo := repository.NewProfileRepo(db)
	profileService := service.NewProfileService(repo)
	profileController := controller.NewProfileController(profileService, client)

	profileHandler := r.Group("/")

	profileHandler.Post("/register/:role", profileController.CreateUser)
	profileHandler.Post("/login/:role", profileController.LoginUser)
	profileHandler.Post("/reset-password", profileController.SendResetPasswordLink)
	profileHandler.Put("/reset-password/:code", profileController.ResetPassword)
}
