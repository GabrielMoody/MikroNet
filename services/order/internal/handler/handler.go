package handler

import (
	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/controller"
	"github.com/GabrielMoody/MikroNet/services/order/internal/events"
	"github.com/GabrielMoody/MikroNet/services/order/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/order/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

func OrderHandler(r fiber.Router, db *gorm.DB, rdb *redis.Client, amqp_cons, amqp_pub *common.AMQP) events.OrderEvents {
	repo := repository.NewOrderRepo(db, rdb)
	service := service.NewOrderService(repo, amqp_pub)
	events := events.NewEvents(service, amqp_cons)
	controller := controller.NewOrderController(service)

	api := r.Group("/")
	api.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"status": "pass"})
	})

	api.Get("/:orderID", controller.GetOrderByID)

	return events
}
