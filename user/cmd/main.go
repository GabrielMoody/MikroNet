package main

import (
	"fmt"
	"log"
	"net"

	"github.com/GabrielMoody/mikronet-user-service/internal/handler"
	"github.com/GabrielMoody/mikronet-user-service/internal/model"
	"github.com/GabrielMoody/mikronet-user-service/internal/pb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	app := fiber.New()
	grpcServer := grpc.NewServer()
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Authorization, Content-Type",
		AllowOrigins: "*",
	}))

	app.Use(logger.New(logger.Config{
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

	db := model.DatabaseInit()

	grpcHandler := handler.GRPCHandler(db)
	pb.RegisterUserServiceServer(grpcServer, grpcHandler)
	reflection.Register(grpcServer)

	api := app.Group("/")

	handler.UserHandler(api, db)

	go func() {
		fmt.Println("gRPC server running on 5005")
		lis, err := net.Listen("tcp", "0.0.0.0:5005")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	err := app.Listen("0.0.0.0:8010")

	if err != nil {
		log.Fatal(err)
	}
}
