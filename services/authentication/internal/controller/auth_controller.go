package controller

import (
	"os"
	"time"

	"github.com/GabrielMoody/mikronet-auth-service/internal/dto"
	"github.com/GabrielMoody/mikronet-auth-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController interface {
	CreateUser(c *fiber.Ctx) error
	CreateDriver(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func (a *AuthControllerImpl) CreateUser(c *fiber.Ctx) error {
	var user dto.UserRegistrationsReq
	ctx := c.Context()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	_, errService := a.AuthService.CreateUserService(ctx, user)

	if errService != nil && errService.ValidationErrors != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": errService.ValidationErrors,
		})
	}

	if errService != nil && errService.Err != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": errService.Err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "akun berhasil dibuat!",
	})
}

func (a *AuthControllerImpl) CreateDriver(c *fiber.Ctx) error {
	var driver dto.DriverRegistrationsReq
	ctx := c.Context()

	if err := c.BodyParser(&driver); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	_, errService := a.AuthService.CreateDriverService(ctx, driver)

	if errService != nil && errService.ValidationErrors != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": errService.ValidationErrors,
		})
	}

	if errService != nil && errService.Err != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": errService.Err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "akun berhasil dibuat!",
	})
}

func (a *AuthControllerImpl) LoginUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var user dto.UserLoginReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	res, errService := a.AuthService.LoginUserService(ctx, user)

	if errService != nil && errService.ValidationErrors != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": errService.ValidationErrors,
		})
	}

	if errService != nil && errService.Err != nil {
		return c.Status(errService.Code).JSON(fiber.Map{
			"status": "error",
			"errors": errService.Err.Error(),
		})
	}

	claims := jwt.MapClaims{
		"id":    res.ID,
		"email": res.Email,
		"role":  res.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   os.Getenv("JWT_ISS"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if errToken != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "error",
			"errors": "Invalid token",
		})
	}

	refreshClaims := jwt.MapClaims{
		"id":  res.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iss": os.Getenv("JWT_ISS"),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, errRefreshToken := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if errRefreshToken != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "error",
			"errors": "Invalid refresh token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"access_token":  t,
			"refresh_token": rt,
		},
	})
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}
