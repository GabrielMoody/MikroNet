package main

import (
	"github.com/GabrielMoody/mikroNet/notification/internal/handler"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	api := app.Group("/")

	handler.NewHandler(api)

	err := app.Listen("0.0.0.0:8015")

	if err != nil {
		log.Fatal(err)
	}
}
