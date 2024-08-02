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

	profileHandler := r.Group("/profile")

	profileHandler.Get("/:userid", profileController.GetUser)
	profileHandler.Post("", profileController.CreateUser)
	profileHandler.Post("/login", profileController.LoginUser)
	profileHandler.Delete("/:userid", profileController.DeleteUser)
	profileHandler.Patch("/:userid", profileController.UpdateUser)
	profileHandler.Patch("/:userid", profileController.ChangePassword)
	profileHandler.Patch("", profileController.ForgotPassword)
}
