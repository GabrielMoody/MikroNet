package main

import (
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/controller"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/logger"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/models"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})

	logger := logger.New()

	app.Use(middleware.LoggerMiddleware(logger))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	db := models.DatabaseInit()
	repo := repository.NewAuthRepo(db)
	authService := service.NewAuthService(repo)
	authController := controller.NewAuthController(authService)

	api := app.Group("/")

	handler.AuthHandler(api, authController)

	err := app.Listen(":8050")
	if err != nil {
		return
	}
}
