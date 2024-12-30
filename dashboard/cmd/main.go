package main

import (
	"github.com/GabrielMoody/mikroNet/dashboard/internal/handler"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/models"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/pb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	app.Use(logger.New())

	db := models.DatabaseInit()

	api := app.Group("/")

	userConn, err := grpc.NewClient(":5005", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer userConn.Close()

	driverConn, err := grpc.NewClient(":5006", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer driverConn.Close()

	driver := pb.NewDriverServiceClient(driverConn)
	user := pb.NewUserServiceClient(userConn)

	handler.DashboardHandler(api, db, driver, user)

	err = app.Listen(":8030")
	if err != nil {
		return
	}
}
