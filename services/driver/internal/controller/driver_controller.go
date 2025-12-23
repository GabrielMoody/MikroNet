package controller

import (
	"os"

	"github.com/GabrielMoody/MikroNet/services/driver/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DriverController interface {
	GetDriver(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
	SetStatus(c *fiber.Ctx) error
	ConfirmOrder(c *fiber.Ctx) error
}

type DriverControllerImpl struct {
	service service.DriverService
}

func (a *DriverControllerImpl) ConfirmOrder(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	var req dto.OrderConfirmation

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	a.service.ConfirmOrder(c.Context(), orderId, req.IsAccepted)

	return c.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}

func (a *DriverControllerImpl) GetDriver(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetDriverDetails(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"id":           res.ID,
			"name":         res.Name,
			"email":        res.Email,
			"plate_number": res.LicenseNumber,
		},
	})
}

func (a *DriverControllerImpl) GetStatus(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetStatus(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"status": res,
		},
	})
}

func (a *DriverControllerImpl) SetStatus(c *fiber.Ctx) error {
	var data dto.StatusReq
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	res, err := a.service.SetStatus(ctx, payload["id"].(string), data)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"status": res,
		},
	})
}

func NewDriverController(service service.DriverService) DriverController {
	return &DriverControllerImpl{
		service: service,
	}
}
