package handler

import (
	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/controller"
	"github.com/GabrielMoody/MikroNet/services/order/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/order/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/order/internal/service"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func UserHandler(r fiber.Router, db *gorm.DB, amqp *common.AMQP) {
	repo := repository.NewUserRepo(db)
	serviceUser := service.NewUserService(repo, amqp)
	controllerUser := controller.NewUserController(serviceUser)

	api := r.Group("/")

	api.Use(middleware.ValidateUserRole)

	api.Get("/", controllerUser.GetUser)

	api.Post("/order", controllerUser.Order)
}
