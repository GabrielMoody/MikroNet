package handler

import (
	"github.com/GabrielMoody/mikroNet/dashboard/internal/controller"
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

	api.Get("/users", controllerDashboard.GetUsers)
	api.Get("/users/:id", controllerDashboard.GetUserDetails)

	api.Get("/drivers", controllerDashboard.GetDrivers)
	api.Get("/drivers/:id", controllerDashboard.GetDriverDetails)

	api.Get("/owners", controllerDashboard.GetBusinessOwners)
	api.Get("/owners/:id", controllerDashboard.GetBusinessOwnerDetails)
	api.Get("/owners/blocked", controllerDashboard.GetBlockedBusinessOwners)

	api.Post("/block/:accountId", controllerDashboard.BlockAccount)

	api.Post("/", controllerDashboard.RegisterBusinessOwner)

	api.Get("/reviews", controllerDashboard.GetReviews)
	api.Get("/reviews/:id", controllerDashboard.GetReviewByID)
}
