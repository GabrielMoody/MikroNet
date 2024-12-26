package controller

import (
	"github.com/GabrielMoody/mikroNet/driver/internal/dto"
	"github.com/GabrielMoody/mikroNet/driver/internal/middleware"
	"github.com/GabrielMoody/mikroNet/driver/internal/service"
	"github.com/gofiber/fiber/v2"
	"os"
)

type DriverController interface {
	GetDriver(c *fiber.Ctx) error
	EditDriver(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
	SetStatus(c *fiber.Ctx) error
	GetRequest(c *fiber.Ctx) error
	AcceptRequest(c *fiber.Ctx) error
	GetTripHistories(c *fiber.Ctx) error
	GetAvailableSeats(c *fiber.Ctx) error
	SetAvailableSeats(c *fiber.Ctx) error
}

type DriverControllerImpl struct {
	service service.DriverService
}

func (a *DriverControllerImpl) GetDriver(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetDriverDetails(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) EditDriver(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	var data dto.EditDriverReq

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	res, err := a.service.EditDriverDetails(ctx, payload["id"].(string), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) GetTripHistories(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetTripHistories(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) GetAvailableSeats(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetAvailableSeats(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) SetAvailableSeats(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	var data dto.SeatReq

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	res, err := a.service.SetAvailableSeats(ctx, data, payload["id"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Error",
		"data": fiber.Map{
			"seat": res,
		},
	})
}

func (a *DriverControllerImpl) GetStatus(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetStatus(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) SetStatus(c *fiber.Ctx) error {
	var data dto.StatusReq
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	res, err := a.service.SetStatus(ctx, payload["id"].(string), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DriverControllerImpl) GetRequest(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *DriverControllerImpl) AcceptRequest(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewDriverController(service service.DriverService) DriverController {
	return &DriverControllerImpl{
		service: service,
	}
}
