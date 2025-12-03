package main

import (
	"github.com/GabrielMoody/MikroNet/services/order/config/rabbitmq"
	"github.com/GabrielMoody/MikroNet/services/order/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Authorization, Content-Type",
		AllowOrigins: "*",
	}))
	app.Use(logger.New(logger.Config{
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

	db := model.DatabaseInit()
	rdb := model.RedisConnect()
	aqmp := rabbitmq.Init("amqp://admin:admin123@rabbitmq:5672/")

	handler.OrderHandler(db, rdb, aqmp)
}
