package handler

import (
	"github.com/GabrielMoody/mikroNet/user/internal/controller"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"github.com/GabrielMoody/mikroNet/user/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DriverHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewDriverRepo(db)
	serviceDriver := service.NewDriverService(repo)
	controllerDriver := controller.NewDriverController(serviceDriver)

	api := r.Group("/driver")

	api.Get("/status/:id", controllerDriver.GetStatus)
	api.Post("/status/:id", controllerDriver.SetStatus)
}
