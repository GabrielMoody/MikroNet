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
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	app.Use(logger.New())

	db := models.DatabaseInit()

	api := app.Group("/")

	conn, err := grpc.NewClient(":5005", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	driver := pb.NewDriverServiceClient(conn)
	user := pb.NewUserServiceClient(conn)

	handler.OwnerHandler(api, db, driver, user)

	err = app.Listen("0.0.0.0:8030")
	if err != nil {
		return
	}
}
