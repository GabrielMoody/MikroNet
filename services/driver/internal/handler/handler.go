package handler

import (
	"github.com/GabrielMoody/mikronet-driver-service/internal/controller"
	"github.com/GabrielMoody/mikronet-driver-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-driver-service/internal/repository"
	"github.com/GabrielMoody/mikronet-driver-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DriverHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewDriverRepo(db)
	serviceDriver := service.NewDriverService(repo)
	controllerDriver := controller.NewDriverController(serviceDriver)

	api := r.Group("/")

	api.Get("/images/:id", controllerDriver.GetImage)
	api.Get("/online", controllerDriver.GetAllDriverLastSeen)

	api.Use(middleware.ValidateDriverRole)

	api.Get("/", controllerDriver.GetDriver)
	api.Put("/", controllerDriver.EditDriver)
	api.Get("/status/", controllerDriver.GetStatus)
	api.Put("/status/", controllerDriver.SetStatus)
	api.Get("/code/", controllerDriver.GetQrisData)
	api.Get("/trips/", controllerDriver.GetTripHistories)

	api.Post("/heartbeat/", controllerDriver.SetDriverLastSeen)
}
