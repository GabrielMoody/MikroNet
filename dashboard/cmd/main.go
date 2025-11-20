package main

import (
	"os"

	"github.com/GabrielMoody/mikronet-dashboard-service/config"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/handler"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func main() {
	config.InitLogger()
	defer config.Logger.Sync()

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		ErrorHandler:  config.ErrorHandler,
	})

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()
		status := c.Response().StatusCode()

		if err == nil {
			config.Logger.Info("Incoming request",
				zap.String("service", os.Getenv("SERVICE_NAME")),
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.String("ip", c.IP()),
				zap.Int("status", status),
			)
		}

		return err
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "deny")
		c.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	db := models.DatabaseInit()

	api := app.Group("/")

	handler.DashboardHandler(api, db)

	err := app.Listen("0.0.0.0:8030")
	if err != nil {
		return
	}
}
