package main

import (
	"log"

	"github.com/GabrielMoody/MikroNet/services/user/config/rabbitmq"
	"github.com/GabrielMoody/MikroNet/services/user/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/user/internal/logger"
	"github.com/GabrielMoody/MikroNet/services/user/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/user/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Authorization, Content-Type",
		AllowOrigins: "*",
	}))

	logger := logger.New()

	app.Use(middleware.LoggerMiddleware(logger))

	db := model.DatabaseInit()
	amqp := rabbitmq.Init("amqp://admin:admin123@localhost:15672")

	api := app.Group("/")

	handler.UserHandler(api, db, amqp)

	err := app.Listen("0.0.0.0:8010")

	if err != nil {
		log.Fatal(err)
	}
}
