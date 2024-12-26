package controller

import (
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/middleware"
	"github.com/GabrielMoody/mikroNet/user/internal/service"
	"github.com/gofiber/fiber/v2"
	"os"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	EditUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetRoutes(c *fiber.Ctx) error
	OrderMikro(c *fiber.Ctx) error
	CarterMikro(c *fiber.Ctx) error
	GetTripHistories(c *fiber.Ctx) error
	ReviewOrder(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	service service.UserService
}

func (a *UserControllerImpl) GetUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))

	ctx := c.Context()

	res, err := a.service.GetUserDetails(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"first_name":    res.FirstName,
			"last_name":     res.LastName,
			"email":         res.Email,
			"date_of_birth": res.DateOfBirth,
			"Age":           res.Age,
			"Gender":        res.Gender,
		},
	})
}

func (a *UserControllerImpl) EditUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	var data dto.EditUserDetails

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.service.EditUserDetails(ctx, payload["id"].(string), data)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"first_name":    res.FirstName,
			"last_name":     res.LastName,
			"date_of_birth": res.DateOfBirth,
			"Age":           res.Age,
			"Gender":        res.Gender,
		},
	})
}

func (a *UserControllerImpl) DeleteUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	err := a.service.DeleteUserDetails(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   nil,
	})
}

func (a *UserControllerImpl) ReviewOrder(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var data dto.ReviewReq

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.service.ReviewOrder(ctx, id, data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (a *UserControllerImpl) GetRoutes(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.service.GetRoutes(ctx)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"routes": res,
	})
}

func (a *UserControllerImpl) OrderMikro(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var loc dto.CurrLocation

	if err := c.BodyParser(&loc); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.service.OrderMikro(ctx, loc.Lat, loc.Lon, id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (a *UserControllerImpl) CarterMikro(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *UserControllerImpl) GetTripHistories(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.service.GetTripHistories(ctx, id)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"histories": res,
	})
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}
