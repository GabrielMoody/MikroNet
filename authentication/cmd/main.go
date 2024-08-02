package main

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	db := models.DatabaseInit()

	api := app.Group("/api")

	handler.ProfileHandler(api, db)

	err := app.Listen("0.0.0.0:8000")
	if err != nil {
		return
	}
}
