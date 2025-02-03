package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GabrielMoody/mikronet-driver-service/internal/dto"
	"github.com/GabrielMoody/mikronet-driver-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-driver-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DriverController interface {
	GetDriver(c *fiber.Ctx) error
	EditDriver(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
	SetStatus(c *fiber.Ctx) error
	GetTripHistories(c *fiber.Ctx) error
	GetImage(c *fiber.Ctx) error
}

type DriverControllerImpl struct {
	service service.DriverService
}

func (a *DriverControllerImpl) GetImage(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.service.GetImage(ctx, id)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	var extension = map[string]string{
		"image/jpeg": "jpg",
		"image/png":  "png",
	}

	img, errI := os.ReadFile(res)

	if errI != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": errI,
		})
	}

	ext := http.DetectContentType(img)

	c.Response().Header.Set("Content-Type", fmt.Sprintf("image/%s", extension[ext]))

	return c.Status(fiber.StatusOK).Send(img)
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
			"id":             res.ID,
			"name":           res.Name,
			"email":          res.Email,
			"license_number": res.LicenseNumber,
			"SIM":            res.SIM,
			"image":          os.Getenv("BASE_URL") + "/api/driver/images/" + res.ID,
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
