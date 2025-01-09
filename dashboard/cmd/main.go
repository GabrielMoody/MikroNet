package main

import (
	"fmt"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/handler"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/models"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/pb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})
	grpcServer := grpc.NewServer()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

	db := models.DatabaseInit()

	grpcHandler := handler.GRPCHandler(db)
	pb.RegisterOwnerServiceServer(grpcServer, grpcHandler)
	reflection.Register(grpcServer)

	api := app.Group("/")

	userConn, err := grpc.NewClient(fmt.Sprintf("%s:5005", os.Getenv("GRPC_USER")), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	driverConn, err := grpc.NewClient(fmt.Sprintf("%s:5006", os.Getenv("GRPC_DRIVER")), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer driverConn.Close()

	driver := pb.NewDriverServiceClient(driverConn)
	user := pb.NewUserServiceClient(userConn)

	handler.DashboardHandler(api, db, driver, user)

	go func() {
		fmt.Println("gRPC server running on 5007")
		lis, err := net.Listen("tcp", "0.0.0.0:5007")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	err = app.Listen("0.0.0.0:8030")
	if err != nil {
		return
	}
}
