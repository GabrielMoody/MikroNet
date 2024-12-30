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

	api.Get("/", middleware.CheckHeaderAuthorization, controllerDriver.GetDriver)
	api.Put("/", middleware.CheckHeaderAuthorization, controllerDriver.EditDriver)
	api.Get("/status/", middleware.CheckHeaderAuthorization, controllerDriver.GetStatus)
	api.Put("/status/", middleware.CheckHeaderAuthorization, controllerDriver.SetStatus)
	api.Get("/seats/", middleware.CheckHeaderAuthorization, controllerDriver.GetAvailableSeats)
	api.Put("/seats/", middleware.CheckHeaderAuthorization, controllerDriver.SetAvailableSeats)
	api.Get("/histories/", controllerDriver.GetTripHistories)
}

func GRPCHandler(db *gorm.DB) *gRPC.GRPC {
	repo := repository.NewDriverRepo(db)
	grpc := gRPC.NewgRPC(repo)

	return grpc
}
