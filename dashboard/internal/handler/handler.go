package handler

import (
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/controller"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/repository"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DashboardHandler(r fiber.Router, db *gorm.DB) {
	repo := repository.NewDashboardRepo(db)
	serviceDashboard := service.NewDashboardService(repo)
	controllerDashboard := controller.NewDashboardController(serviceDashboard)

	api := r.Group("/")
	api.Get("/ktp/:id", controllerDashboard.GetKTP)

	api.Get("/users", controllerDashboard.GetUsers)
	api.Get("/users/:id", controllerDashboard.GetUserDetails)
	api.Delete("/users/:id", middleware.ValidateDashboardRole, controllerDashboard.DeleteUser)

	api.Get("/drivers", controllerDashboard.GetDrivers)
	api.Get("/drivers/:id", controllerDashboard.GetDriverDetails)
	api.Post("/drivers/verified/:id", middleware.ValidateDashboardRole, controllerDashboard.SetDriverStatusVerified)
	api.Delete("/drivers/:id", middleware.ValidateDashboardRole, controllerDashboard.DeleteDriver)

	api.Get("/block", middleware.ValidateDashboardRole, controllerDashboard.GetAllBlockAccount)
	api.Post("/block/:id", middleware.ValidateDashboardRole, controllerDashboard.BlockAccount)
	api.Put("/block/:id", middleware.ValidateDashboardRole, controllerDashboard.UnblockAccount)

	api.Get("/reviews", controllerDashboard.GetReviews)
	api.Get("/reviews/:id", controllerDashboard.GetReviewByID)

	api.Get("/routes", controllerDashboard.GetRoutes)
	api.Post("/route", middleware.ValidateDashboardRole, controllerDashboard.AddRoute)
	api.Put("/route/:id", middleware.ValidateDashboardRole, controllerDashboard.EditAmountRoute)
	api.Delete("/route/:id", middleware.ValidateDashboardRole, controllerDashboard.DeleteRoute)

	api.Get("/histories", controllerDashboard.GetAllTripHistories)

	api.Get("/reports", controllerDashboard.MonthlyReport)
}
