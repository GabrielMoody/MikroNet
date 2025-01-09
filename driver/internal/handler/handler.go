package handler

import (
	"github.com/GabrielMoody/mikroNet/driver/internal/controller"
	"github.com/GabrielMoody/mikroNet/driver/internal/gRPC"
	"github.com/GabrielMoody/mikroNet/driver/internal/middleware"
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

	api.Use(middleware.ValidateDriverRole)

	api.Get("/", controllerDriver.GetDriver)
	api.Put("/", controllerDriver.EditDriver)
	api.Get("/status/", controllerDriver.GetStatus)
	api.Put("/status/", controllerDriver.SetStatus)
	api.Get("/seats/", controllerDriver.GetAvailableSeats)
	api.Put("/seats/", controllerDriver.SetAvailableSeats)
}

func GRPCHandler(db *gorm.DB) *gRPC.GRPC {
	repo := repository.NewDriverRepo(db)
	grpc := gRPC.NewgRPC(repo)

	return grpc
}
