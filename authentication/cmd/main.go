package main

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/GabrielMoody/MikroNet/authentication/internal/pb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	userConn, err := grpc.NewClient(":5005", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	driverConn, err := grpc.NewClient(":5006", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer driverConn.Close()
	defer userConn.Close()

	userPB := pb.NewUserServiceClient(userConn)
	driverPB := pb.NewDriverServiceClient(driverConn)

	db := models.DatabaseInit()

	api := app.Group("/")

	handler.ProfileHandler(api, db, userPB, driverPB)

	err = app.Listen(":8000")
	if err != nil {
		return
	}
}
