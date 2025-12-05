package handler

import (
	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/mikronet-user-service/internal/controller"
	"github.com/GabrielMoody/mikronet-user-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-user-service/internal/repository"
	"github.com/GabrielMoody/mikronet-user-service/internal/service"
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

	api.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{"status": "pass"})
	})
}
