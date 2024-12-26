package handler

import (
	"github.com/GabrielMoody/mikroNet/dashboard/internal/controller"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/pb"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func OwnerHandler(r fiber.Router, db *gorm.DB, driver pb.DriverServiceClient, user pb.UserServiceClient) {
	//repo := repository.NewOwnerRepo(db)
	//serviceOwner := service.NewOwnerService(repo)
	controllerOwner := controller.NewOwnerController(driver, user)

	api := r.Group("/")

	api.Post("/", controllerOwner.RegisterBusinessOwner)
	api.Post("/driver", controllerOwner.RegisterNewDriver)
	api.Get("/users", controllerOwner.GetUsers)
	api.Get("/ratings/:id", controllerOwner.GetRatings)
	api.Get("/status/:id", controllerOwner.GetStatus)
}
