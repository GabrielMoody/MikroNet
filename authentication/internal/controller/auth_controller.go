package controller

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type ProfileController interface {
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	SendResetPasswordLink(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
}

type ProfileControllerImpl struct {
	ProfileService service.ProfileService
}

func (a *ProfileControllerImpl) CreateUser(c *fiber.Ctx) error {
	User := new(dto.UserRegistrationsReq)
	Ctx := c.Context()
	role := c.Params("role")

	if (role != "user") && (role != "driver") {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "No such request",
		})
	}

	if err := c.BodyParser(&User); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.CreateUserService(Ctx, *User, role)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
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
	role := c.Params("role")
	User := new(dto.UserLoginReq)

	if !(role == "user" || role == "driver" || role == "admin" || role == "government" || role == "business_owner") {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
		})
	}

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
		"role":  res.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iss":   os.Getenv("JWT_ISS"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, errToken := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if errToken != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  t,
	})
}

func (a *ProfileControllerImpl) SendResetPasswordLink(c *fiber.Ctx) error {
	var email dto.ForgotPasswordReq
	ctx := c.Context()

	if err := c.BodyParser(&email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.SendResetPasswordService(ctx, email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"err":    err.Err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *ProfileControllerImpl) ResetPassword(c *fiber.Ctx) error {
	code := c.Params("code")
	var rp dto.ResetPasswordReq
	ctx := c.Context()

	if err := c.BodyParser(&rp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.ProfileService.ResetPassword(ctx, rp, code)

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

func NewProfileController(profileService service.ProfileService) ProfileController {
	return &ProfileControllerImpl{ProfileService: profileService}
}
