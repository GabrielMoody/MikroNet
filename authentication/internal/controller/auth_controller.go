package controller

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/middleware"
	"github.com/GabrielMoody/MikroNet/authentication/internal/pb"
	"github.com/GabrielMoody/MikroNet/authentication/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthController interface {
	CreateUser(c *fiber.Ctx) error
	CreateDriver(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	SendResetPasswordLink(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
}

type AuthControllerImpl struct {
	AuthService service.AuthService
	pbUser      pb.UserServiceClient
	pbDriver    pb.DriverServiceClient
}

func (a *AuthControllerImpl) ChangePassword(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))

	Ctx := c.Context()
	var user dto.ChangePasswordReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.AuthService.ChangePasswordService(Ctx, payload["id"].(string), user)

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

func (a *AuthControllerImpl) CreateUser(c *fiber.Ctx) error {
	var user dto.UserRegistrationsReq
	ctx := c.Context()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.AuthService.CreateUserService(ctx, user, "user")

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	_, errPb := a.pbUser.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Id:          res,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Password:    user.Password,
			DateOfBirth: user.DateOfBirth,
			Age:         uint32(user.Age),
		},
	})

	if errPb != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  errPb.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *AuthControllerImpl) CreateDriver(c *fiber.Ctx) error {
	var driver dto.DriverRegistrationsReq
	ctx := c.Context()

	if err := c.BodyParser(&driver); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.AuthService.CreateDriverService(ctx, driver, "driver")

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"error":  err.Err.Error(),
		})
	}

	_, errPb := a.pbDriver.CreateDriver(ctx, &pb.CreateDriverRequest{
		Id:            res,
		FirstName:     driver.FirstName,
		LastName:      driver.LastName,
		Email:         driver.Email,
		Password:      driver.Password,
		Age:           uint32(driver.Age),
		PhoneNumber:   driver.PhoneNumber,
		DateOfBirth:   driver.DateOfBirth,
		LicenseNumber: driver.LicenseNumber,
	})

	if errPb != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  errPb.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   res,
	})
}

func (a *AuthControllerImpl) LoginUser(c *fiber.Ctx) error {
	Ctx := c.Context()
	var user dto.UserLoginReq

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.AuthService.LoginUserService(Ctx, user)

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
		"data": fiber.Map{
			"token": t,
		},
	})
}

func (a *AuthControllerImpl) SendResetPasswordLink(c *fiber.Ctx) error {
	var email dto.ForgotPasswordReq
	ctx := c.Context()

	if err := c.BodyParser(&email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.AuthService.SendResetPasswordService(ctx, email)

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

func (a *AuthControllerImpl) ResetPassword(c *fiber.Ctx) error {
	code := c.Params("code")
	var rp dto.ResetPasswordReq
	ctx := c.Context()

	if err := c.BodyParser(&rp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	res, err := a.AuthService.ResetPassword(ctx, rp, code)

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

func NewAuthController(authService service.AuthService, user pb.UserServiceClient, driver pb.DriverServiceClient) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
		pbUser:      user,
		pbDriver:    driver,
	}
}
