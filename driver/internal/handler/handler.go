package handler

import (
	"github.com/GabrielMoody/mikroNet/driver/internal/controller"
	"github.com/GabrielMoody/mikroNet/driver/internal/repository"
	"github.com/GabrielMoody/mikroNet/driver/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DriverHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewDriverRepo(db)
	serviceDriver := service.NewDriverService(repo)
	controllerDriver := controller.NewDriverController(serviceDriver)

	api := r.Group("/")

	api.Get("/driver/", controllerDriver.GetDriver)
	api.Patch("/driver/", controllerDriver.EditDriver)
	api.Get("/status/", controllerDriver.GetStatus)
	api.Patch("/status/", controllerDriver.SetStatus)
	api.Get("/seats/", controllerDriver.GetAvailableSeats)
	api.Patch("/seats/", controllerDriver.SetAvailableSeats)
	api.Get("/histories/", controllerDriver.GetTripHistories)
}
