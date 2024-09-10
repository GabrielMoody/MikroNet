package handler

import (
	"github.com/GabrielMoody/mikroNet/business_owner/internal/controller"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/repository"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OwnerHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewOwnerRepo(db)
	serviceOwner := service.NewOwnerService(repo)
	controllerOwner := controller.NewOwnerController(serviceOwner)

	api := r.Group("/owner")

	api.Post("/", controllerOwner.RegisterBusinessOwner)
	api.Post("/driver", controllerOwner.RegisterNewDriver)
	api.Get("/drivers/:id", controllerOwner.GetDrivers)
	api.Get("/ratings/:id", controllerOwner.GetRatings)
	api.Get("/status/:id", controllerOwner.GetStatus)
}
