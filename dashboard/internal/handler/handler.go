package handler

import (
	"github.com/GabrielMoody/mikroNet/dashboard/internal/controller"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/gRPC"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/middleware"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/pb"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/repository"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DashboardHandler(r fiber.Router, db *gorm.DB, driver pb.DriverServiceClient, user pb.UserServiceClient) {
	repo := repository.NewDashboardRepo(db)
	serviceDashboard := service.NewDashboardService(repo)
	controllerDashboard := controller.NewDashboardController(serviceDashboard, driver, user)

	api := r.Group("/")

	api.Use(middleware.ValidateDashboardRole)

	api.Get("/users", controllerDashboard.GetUsers)
	api.Get("/users/:id", controllerDashboard.GetUserDetails)

	api.Get("/drivers", controllerDashboard.GetDrivers)
	api.Get("/drivers/:id", controllerDashboard.GetDriverDetails)

	api.Get("/owners", controllerDashboard.GetBusinessOwners)
	api.Get("/owners/blocked", controllerDashboard.GetBlockedBusinessOwners)
	api.Get("/owners/unverified", controllerDashboard.GetUnverifiedBusinessOwners)
	api.Get("/owners/:id", controllerDashboard.GetBusinessOwnerDetails)

	api.Put("/owners/verified/:id", controllerDashboard.SetStatusVerified)

	api.Post("/block/:id", controllerDashboard.BlockAccount)
	api.Delete("/block/:id", controllerDashboard.UnblockAccount)

	api.Get("/reviews", controllerDashboard.GetReviews)
	api.Get("/reviews/:id", controllerDashboard.GetReviewByID)
}

func GRPCHandler(db *gorm.DB) *gRPC.GRPC {
	repo := repository.NewDashboardRepo(db)
	grpc := gRPC.NewGRPC(repo)

	return grpc
}
