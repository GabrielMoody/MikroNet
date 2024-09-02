package controller

import (
	"github.com/GabrielMoody/mikroNet/profiles/internal/dto"
	"github.com/GabrielMoody/mikroNet/profiles/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ProfileController interface {
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
	SendResetPassword(c *fiber.Ctx) error
}

type ProfileControllerImpl struct {
	ProfileService service.ProfileService
}

func (a *ProfileControllerImpl) UpdateUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	user := new(dto.UserChangeProfileReq)
	id := c.Params("id")

	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.EditUserService(Ctx, id, *user)

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

func (a *ProfileControllerImpl) DeleteUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	id := c.Params("id")
	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	res, err := a.ProfileService.DeleteUserService(Ctx, id)

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

func (a *ProfileControllerImpl) GetUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	id := c.Params("id")

	//jwtUser := c.Locals("user").(*jwt.Token)
	//claims := jwtUser.Claims.(jwt.MapClaims)
	//id := claims["id"].(string)

	res, err := a.ProfileService.GetUserService(Ctx, id)

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

func (a *ProfileControllerImpl) ChangePassword(c *fiber.Ctx) error {
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

	res, err := a.ProfileService.ChangePasswordService(Ctx, id, user)

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

func (a *ProfileControllerImpl) ForgotPassword(c *fiber.Ctx) error {
	var email dto.ForgotPasswordReq

	if err := c.BodyParser(&email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return nil
}

func NewProfileController(profileService service.ProfileService) ProfileController {
	return &ProfileControllerImpl{ProfileService: profileService}
}
