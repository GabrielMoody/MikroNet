package main

import (
	"github.com/GabrielMoody/MikroNet/services/driver/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/logger"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

	api := app.Group("/")

	handler.DriverHandler(api, db)

	err := app.Listen("0.0.0.0:8020")

	if err != nil {
		log.Fatal(err)
	}
}
