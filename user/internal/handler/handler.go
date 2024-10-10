package handler

import (
	"github.com/GabrielMoody/mikroNet/user/internal/controller"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"github.com/GabrielMoody/mikroNet/user/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewUserRepo(db)
	serviceUser := service.NewUserService(repo)
	controllerUser := controller.NewUserController(serviceUser)

	api := r.Group("/")

	api.Get("/histories/:id", controllerUser.GetTripHistories)
	api.Get("/routes", controllerUser.GetRoutes)
	api.Post("/orders/:id", controllerUser.OrderMikro)
}
