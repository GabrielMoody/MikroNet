package main

import (
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/handler"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/model"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

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

	err := app.Listen(":8040")

	if err != nil {
		log.Fatal(err)
	}
}
