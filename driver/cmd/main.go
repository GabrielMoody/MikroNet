package main

import (
	"github.com/GabrielMoody/mikroNet/driver/internal/handler"
	"github.com/GabrielMoody/mikroNet/driver/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	db := model.DatabaseInit()

	api := app.Group("/api")

	handler.DriverHandler(api, db)

	err := app.Listen("0.0.0.0:8013")

	if err != nil {
		log.Fatal(err)
	}
}
