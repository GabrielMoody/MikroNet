package controller

import (
	"net/http"
	"os"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/dto"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DashboardController interface {
	SetDriverStatusVerified(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUserDetails(c *fiber.Ctx) error
	GetDrivers(c *fiber.Ctx) error
	GetDriverDetails(c *fiber.Ctx) error
	GetAllTripHistories(c *fiber.Ctx) error
	EditAmountRoute(c *fiber.Ctx) error
	DeleteDriver(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	BlockAccount(c *fiber.Ctx) error
	UnblockAccount(c *fiber.Ctx) error
	GetReviews(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
	GetAllBlockAccount(c *fiber.Ctx) error
	AddRoute(c *fiber.Ctx) error
	MonthlyReport(c *fiber.Ctx) error
	GetKTP(c *fiber.Ctx) error
	GetRoutes(c *fiber.Ctx) error
	DeleteRoute(c *fiber.Ctx) error
}

type DashboardControllerImpl struct {
	DashboardService service.DashboardService
}

func (a *DashboardControllerImpl) DeleteRoute(c *fiber.Ctx) error {
	ctx := c.Context()

	id := c.Params("id")

	res, err := a.DashboardService.DeleteRoute(ctx, id)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetRoutes(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetRoutes(ctx)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetKTP(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.GetImage(ctx, id)

	if err != nil {
		return err
	}

	img, errI := os.ReadFile(res)

	if errI != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": errI,
		})
	}

	ext := http.DetectContentType(img)

	c.Response().Header.Set("Content-Type", ext)

	return c.Status(fiber.StatusOK).Send(img)
}

func (a *DashboardControllerImpl) GetAllTripHistories(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetAllHistories(ctx)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) EditAmountRoute(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	var b dto.EditAmount
	if err := c.BodyParser(&b); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	_, err := a.DashboardService.EditAmountRoute(ctx, b, id)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   "Berhasil memperbaharui harga!",
	})
}

func (a *DashboardControllerImpl) MonthlyReport(c *fiber.Ctx) error {
	ctx := c.Context()

	var q dto.MonthReport
	if err := c.QueryParser(&q); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	res, err := a.DashboardService.MonthlyReport(ctx, q)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) AddRoute(c *fiber.Ctx) error {
	ctx := c.Context()

	var body dto.AddRoute
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	res, err := a.DashboardService.AddRoute(ctx, body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetAllBlockAccount(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetAllBlockAccount(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"block_accounts": res,
			"count":          len(res),
		},
	})
}

func (a *DashboardControllerImpl) DeleteDriver(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	_, err := a.DashboardService.DeleteDriver(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Berhasil menghapus akun!",
	})
}

func (a *DashboardControllerImpl) DeleteUser(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	_, err := a.DashboardService.DeleteUser(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Berhasil menghapus akun!",
	})
}

func (a *DashboardControllerImpl) SetDriverStatusVerified(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.SetDriverStatusVerified(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) UnblockAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	accountId := c.Params("id")

	_, err := a.DashboardService.UnblockAccount(ctx, accountId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   "Berhasil membuka blokir akun!",
	})
}

func (a *DashboardControllerImpl) GetReviews(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetAllReviews(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"reviews": res,
			"count":   len(res),
		},
	})
}

func (a *DashboardControllerImpl) GetReviewByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.GetReviewById(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) BlockAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	accountId := c.Params("id")

	_, err := a.DashboardService.BlockAccount(ctx, accountId)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Berhasil memblokir akun!",
	})
}

func (a *DashboardControllerImpl) GetDriverDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.DashboardService.GetDriverById(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetDrivers(c *fiber.Ctx) error {
	var q dto.GetDriverQuery
	ctx := c.Context()

	if err := c.QueryParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	res, err := a.DashboardService.GetAllDrivers(ctx, q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"drivers": res,
			"count":   len(res),
		},
	})
}

func (a *DashboardControllerImpl) GetUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	res, err := a.DashboardService.GetAllPassengers(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"users": res,
			"count": len(res),
		},
	})
}

func (a *DashboardControllerImpl) GetUserDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()

	res, err := a.DashboardService.GetPassengerById(ctx, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func NewDashboardController(service service.DashboardService) DashboardController {
	return &DashboardControllerImpl{DashboardService: service}
}
