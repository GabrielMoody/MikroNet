package main

import (
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	app.Use(logger.New(logger.Config{
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

	db := models.DatabaseInit()

	api := app.Group("/")

	handler.AuthHandler(api, db)

	err := app.Listen(":8050")
	if err != nil {
		return
	}
}
