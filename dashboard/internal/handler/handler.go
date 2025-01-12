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

	api.Get("/users", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetUsers)
	api.Get("/users/:id", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetUserDetails)

	api.Get("/drivers", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetDrivers)
	api.Get("/drivers/:id", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetDriverDetails)
	api.Put("/drivers/verified/:id", middleware.ValidateDashboardRole("admin"), controllerDashboard.SetDriverStatusVerified)

	api.Get("/owners", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetBusinessOwners)
	api.Get("/owners/blocked", middleware.ValidateDashboardRole("admin", "government"), controllerDashboard.GetBlockedBusinessOwners)
	api.Get("/owners/unverified", middleware.ValidateDashboardRole("admin", "government"), controllerDashboard.GetUnverifiedBusinessOwners)
	api.Get("/owners/:id", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetBusinessOwnerDetails)
	api.Put("/owners/verified/:id", middleware.ValidateDashboardRole("admin"), controllerDashboard.SetOwnerStatusVerified)

	api.Post("/block/:id", middleware.ValidateDashboardRole("admin"), controllerDashboard.BlockAccount)
	api.Delete("/block/:id", middleware.ValidateDashboardRole("admin"), controllerDashboard.UnblockAccount)

	api.Get("/reviews", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetReviews)
	api.Get("/reviews/:id", middleware.ValidateDashboardRole("admin", "owner", "government"), controllerDashboard.GetReviewByID)
}

func GRPCHandler(db *gorm.DB) *gRPC.GRPC {
	repo := repository.NewDashboardRepo(db)
	grpc := gRPC.NewGRPC(repo)

	return grpc
}
