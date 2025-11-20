package controller

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"

	"github.com/GabrielMoody/mikronet-user-service/internal/dto"
	"github.com/GabrielMoody/mikronet-user-service/internal/middleware"
	"github.com/GabrielMoody/mikronet-user-service/internal/service"
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
	Order(c *fiber.Ctx) error
	ReviewOrder(c *fiber.Ctx) error
	Transaction(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	service service.UserService
}

func (a *UserControllerImpl) Transaction(c *fiber.Ctx) error {
	ctx := c.Context()
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token, os.Getenv("JWT_SECRET"))

	var data dto.Transaction

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Error",
			"errors": err.Error(),
		})
	}

	_, err := a.service.Transaction(ctx, data, payload["id"].(string))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   "Berhasil melakukan transaksi",
	})
}

func (a *UserControllerImpl) Order(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token, os.Getenv("JWT_SECRET"))

	var data dto.MessageLoc
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	headers := http.Header{}
	headers.Add("Authorization", token)

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s/tracking/ws/location", os.Getenv("BASE_URL")), headers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": "Failed to connect to WebSocket",
		})
	}

	defer conn.Close()

	location := map[string]interface{}{
		"user_id": payload["id"].(string),
		"role":    payload["role"].(string),
		"lat":     data.Lat,
		"lng":     data.Lng,
	}

	locationJSON, err := json.Marshal(location)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": "Failed to marshal location to JSON",
		})
	}

	if err := conn.WriteMessage(websocket.TextMessage, locationJSON); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": "Failed to send location",
		})
	}

	// Close the WebSocket connection
	if err := conn.Close(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": "Failed to close WebSocket connection",
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
			"id":    res.ID,
			"email": res.Email,
			"name":  res.Name,
		},
	})
}

func (a *UserControllerImpl) ReviewOrder(c *fiber.Ctx) error {
	ctx := c.Context()
	driverId := c.Params("driverId")

	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token, os.Getenv("JWT_SECRET"))

	var data dto.ReviewReq

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Error",
			"errors": err.Error(),
		})
	}

	res, err := a.service.ReviewOrder(ctx, data, payload["id"].(string), driverId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}
