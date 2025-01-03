package main

import (
	"fmt"
	"github.com/GabrielMoody/MikroNet/authentication/internal/handler"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/GabrielMoody/MikroNet/authentication/internal/pb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	userConn, err := grpc.NewClient(fmt.Sprintf("%s:5005", os.Getenv("GRPC_USER")), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	driverConn, err := grpc.NewClient(fmt.Sprintf("%s:5006", os.Getenv("GRPC_DRIVER")), grpc.WithTransportCredentials(insecure.NewCredentials()))

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

	err = app.Listen(":8050")
	if err != nil {
		return
	}
}
