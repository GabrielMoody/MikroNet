package controller

import (
	"strconv"

	"github.com/GabrielMoody/MikroNet/services/order/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	GetOrderByID(c *fiber.Ctx) error
}

type OrderControllerImpl struct {
	service service.OrderService
}

func (a *OrderControllerImpl) GetOrderByID(c *fiber.Ctx) error {
	orderId := c.Params("orderID")

	id, _ := strconv.Atoi(orderId)

	res, err := a.service.GetOrderByID(c.Context(), id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": res,
	})
}

func NewOrderController(service service.OrderService) OrderController {
	return &OrderControllerImpl{
		service: service,
	}
}
