package handler

import (
	"github.com/GabrielMoody/mikroNet/profiles/internal/controller"
	"github.com/GabrielMoody/mikroNet/profiles/internal/repository"
	"github.com/GabrielMoody/mikroNet/profiles/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProfileHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewProfileRepo(db)
	serviceProfile := service.NewProfileService(repo)
	controllerProfile := controller.NewProfileController(serviceProfile)

	api := r.Group("/profile")

	api.Get("/:id", controllerProfile.GetUser)
	api.Put("/:id", controllerProfile.UpdateUser)
	api.Delete("/:id", controllerProfile.DeleteUser)
	api.Post("/:id", controllerProfile.ChangePassword)
}
