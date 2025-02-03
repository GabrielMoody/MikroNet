package handler

import (
	"github.com/GabrielMoody/mikronet-user-service/internal/controller"
	"github.com/GabrielMoody/mikronet-user-service/internal/gRPC"
	"github.com/GabrielMoody/mikronet-user-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-user-service/internal/repository"
	"github.com/GabrielMoody/mikronet-user-service/internal/service"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func UserHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewUserRepo(db)
	serviceUser := service.NewUserService(repo)
	controllerUser := controller.NewUserController(serviceUser)

	api := r.Group("/")

	api.Use(middleware.ValidateUserRole)

	api.Get("/", controllerUser.GetUser)

	api.Post("/order", controllerUser.Order)
	api.Post("/review/:driverId", controllerUser.ReviewOrder)
}

func GRPCHandler(db *gorm.DB) *gRPC.GRPC {
	repo := repository.NewUserRepo(db)
	grpc := gRPC.NewgRPC(repo)

	return grpc
}
