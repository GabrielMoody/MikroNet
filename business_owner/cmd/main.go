package main

import (
	"github.com/GabrielMoody/mikroNet/business_owner/internal/handler"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	db := models.DatabaseInit()

	api := app.Group("/")

	handler.OwnerHandler(api, db)

	err := app.Listen("0.0.0.0:8012")
	if err != nil {
		return
	}
}
