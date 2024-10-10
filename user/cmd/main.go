package main

import (
	"github.com/GabrielMoody/mikroNet/user/internal/handler"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	db := model.DatabaseInit()

	api := app.Group("/")

	handler.UserHandler(api, db)

	err := app.Listen("0.0.0.0:8014")

	if err != nil {
		log.Fatal(err)
	}
}
