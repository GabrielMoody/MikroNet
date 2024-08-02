package controller

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	"github.com/GabrielMoody/MikroNet/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ProfileController interface {
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
}

type ProfileControllerImpl struct {
	ProfileService service.ProfileService
}

func (a *ProfileControllerImpl) CreateUser(c *fiber.Ctx) error {
	User := new(dto.UserRegistrationsReq)
	Ctx := c.Context()

	if err := c.BodyParser(&User); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.CreateUserService(Ctx, *User)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *ProfileControllerImpl) LoginUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	User := new(dto.UserLoginReq)

	if err := c.BodyParser(&User); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.LoginUserService(Ctx, *User)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	claims := jwt.MapClaims{
		"id":    res.ID,
		"email": res.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	v := helper.LoadEnv()
	t, errToken := token.SignedString([]byte(v.GetString("JWT_SECRET")))

	if errToken != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
		"token":  t,
	})
}

func (a *ProfileControllerImpl) UpdateUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	user := new(dto.UserChangeProfileReq)

	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

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

	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

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

	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

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
	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	id := claims["ID"].(string)

	Ctx := c.Context()
	user := new(dto.ChangePasswordReq)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.ChangePasswordService(Ctx, user.OldPassword, user.NewPassword, id)

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
	//TODO implement me
	panic("implement me")
}

func NewProfileController(profileService service.ProfileService) ProfileController {
	return &ProfileControllerImpl{ProfileService: profileService}
}
