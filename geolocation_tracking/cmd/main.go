package main

import (
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/handler"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/model"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			token := c.GetReqHeaders()["Authorization"]
			c.Locals("token", token)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	api := app.Group("/")
	db := model.DatabaseInit()

	handler.NewWSHandler(api, db)

	err := app.Listen("0.0.0.0:8040")

	if err != nil {
		log.Fatal(err)
	}
}
