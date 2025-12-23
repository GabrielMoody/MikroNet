package controller

import (
	"os"

	"github.com/GabrielMoody/MikroNet/services/user/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/user/internal/middleware"
	"github.com/GabrielMoody/MikroNet/services/user/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	Order(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	service service.UserService
}

func (a *UserControllerImpl) Order(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token, os.Getenv("JWT_SECRET"))

	var order_req dto.OrderReq

	if err := c.BodyParser(&order_req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	id := payload["id"].(float64)

	order_req.UserID = int64(id)

	_, err := a.service.MakeOrder(c.Context(), order_req)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   "berhasil",
	})
}

func (a *UserControllerImpl) GetUser(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token, os.Getenv("JWT_SECRET"))
	ctx := c.Context()

	res, err := a.service.GetUserDetails(ctx, payload["id"].(string))

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err.Err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"id": res.ID,
		},
	})
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}
