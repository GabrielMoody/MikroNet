package main

import (
	"github.com/GabrielMoody/mikronet-driver-service/internal/handler"
	"github.com/GabrielMoody/mikronet-driver-service/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

	api := app.Group("/")

	handler.DriverHandler(api, db)

	err := app.Listen("0.0.0.0:8020")

	if err != nil {
		log.Fatal(err)
	}
}
