package handler

import (
	"github.com/GabrielMoody/MikroNet/services/driver/config/rabbitmq"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/controller"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DriverHandler(r fiber.Router, db *gorm.DB) {
	amqp := rabbitmq.Init("amqp://admin:admin123@rabbitmq:5762/")
	repo := repository.NewDriverRepo(db)
	serviceDriver := service.NewDriverService(repo, amqp)
	controllerDriver := controller.NewDriverController(serviceDriver)

	api := r.Group("/")

	api.Use(middleware.ValidateDriverRole)

	api.Get("/", controllerDriver.GetDriver)
	api.Get("/status/", controllerDriver.GetStatus)
	api.Put("/status/", controllerDriver.SetStatus)
	api.Post("/order/confirm/:orderId", controllerDriver.ConfirmOrder)

	api.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]string{"status": "pass"})
	})
}
