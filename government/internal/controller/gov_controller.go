package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/GabrielMoody/mikroNet/government/internal/dto"
	"github.com/GabrielMoody/mikroNet/government/internal/service"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

type GovController interface {
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}

type GovControllerImpl struct {
	GovService service.GovService
}

func (a *GovControllerImpl) UpdateUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	user := dto.UserChangeGovReq{}
	id := c.Params("id")
	image, _ := c.FormFile("image")

	if image != nil {
		user.Image = fmt.Sprintf("%d.%s", time.Now().Unix(), image.Filename)
	}

	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.GovService.EditUserService(Ctx, id, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	if image != nil {
		if err := c.SaveFile(image, fmt.Sprintf("./uploads/%s", user.Image)); err != nil {
			return err
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *GovControllerImpl) DeleteUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	id := c.Params("id")
	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	res, err := a.GovService.DeleteUserService(Ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *GovControllerImpl) GetUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	id := c.Params("id")

	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	res, err := a.GovService.GetUserService(Ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	var encoded interface{}

	if res.ImageUrl != "" {
		image := fmt.Sprintf("./uploads/%s", res.ImageUrl)
		data, errImage := os.ReadFile(image)

		if errImage != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error",
				"error":  errImage.Error(),
			})
		}

		encoded = base64.StdEncoding.EncodeToString(data)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
		"image":  encoded,
	})
}

func (a *GovControllerImpl) ChangePassword(c *fiber.Ctx) error {
	id := c.Params("id")
	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	Ctx := c.Context()
	var user dto.ChangePasswordReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.GovService.ChangePasswordService(Ctx, id, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func NewGovController(profileService service.GovService) GovController {
	return &GovControllerImpl{GovService: profileService}
}
