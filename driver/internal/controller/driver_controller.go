package controller

import (
	"github.com/GabrielMoody/mikroNet/driver/internal/dto"
	"github.com/GabrielMoody/mikroNet/driver/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DriverController interface {
	GetStatus(c *fiber.Ctx) error
	SetStatus(c *fiber.Ctx) error
	GetRequest(c *fiber.Ctx) error
	AcceptRequest(c *fiber.Ctx) error
}

type DriverControllerImpl struct {
	service service.DriverService
}

func (a *DriverControllerImpl) GetStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.service.GetStatus(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": res,
	})
}

func (a *DriverControllerImpl) SetStatus(c *fiber.Ctx) error {
	var data dto.StatusReq
	id := c.Params("id")
	ctx := c.Context()

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	res, err := a.service.SetStatus(ctx, id, data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": res,
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
