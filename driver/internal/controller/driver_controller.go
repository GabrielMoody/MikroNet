package controller

import (
	"fmt"
	"github.com/GabrielMoody/mikroNet/driver/internal/dto"
	"github.com/GabrielMoody/mikroNet/driver/internal/middleware"
	"github.com/GabrielMoody/mikroNet/driver/internal/service"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
)

type DriverController interface {
	GetDriver(c *fiber.Ctx) error
	EditDriver(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
	SetStatus(c *fiber.Ctx) error
	GetTripHistories(c *fiber.Ctx) error
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
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	image, errI := os.ReadFile(res.ProfilePicture)

	if errI != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": errI,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"id":             res.ID,
			"first_name":     res.FirstName,
			"last_name":      res.LastName,
			"email":          res.Email,
			"license_number": res.RegistrationNumber,
			"age":            res.Age,
			"date_of_birth":  res.DateOfBirth,
			"image": fiber.Map{
				"mime_type": fmt.Sprintf("image/%s", filepath.Ext(res.ProfilePicture)[1:]),
				"data":      image,
			},
		},
	})
}

func (a *DriverControllerImpl) EditDriver(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	var data dto.EditDriverReq

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	res, err := a.service.EditDriverDetails(ctx, payload["id"].(string), data)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
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
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
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
