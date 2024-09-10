package controller

import (
	"github.com/GabrielMoody/mikroNet/business_owner/internal/dto"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OwnerController interface {
	RegisterBusinessOwner(c *fiber.Ctx) error
	GetRatings(c *fiber.Ctx) error
	RegisterNewDriver(c *fiber.Ctx) error
	GetDrivers(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
}

type OwnerControllerImpl struct {
	OwnerService service.OwnerService
}

func (a *OwnerControllerImpl) RegisterBusinessOwner(c *fiber.Ctx) error {
	var req dto.OwnerRegistrationReq
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.OwnerService.RegisterBusinessOwner(ctx, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (a *OwnerControllerImpl) GetRatings(c *fiber.Ctx) error {
	ownerId := c.Params("id")
	ctx := c.Context()

	res, err := a.OwnerService.GetRatings(ctx, ownerId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (a *OwnerControllerImpl) RegisterNewDriver(c *fiber.Ctx) error {
	var req dto.DriverRegistrationReq
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.OwnerService.RegisterNewDriver(ctx, req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (a *OwnerControllerImpl) GetDrivers(c *fiber.Ctx) error {
	ownerId := c.Params("id")
	ctx := c.Context()

	res, err := a.OwnerService.GetDrivers(ctx, ownerId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (a *OwnerControllerImpl) GetStatus(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.OwnerService.GetOwnerStatusVerified(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func NewOwnerController(ownerService service.OwnerService) OwnerController {
	return &OwnerControllerImpl{OwnerService: ownerService}
}
