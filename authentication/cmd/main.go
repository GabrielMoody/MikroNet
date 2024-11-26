package main

import (
	"fmt"
	"github.com/GabrielMoody/MikroNet/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	db := models.DatabaseInit()

	fmt.Println("Success published message")

	api := app.Group("/")

	handler.ProfileHandler(api, db)

	err = app.Listen("0.0.0.0:8010")
	if err != nil {
		return
	}
}
