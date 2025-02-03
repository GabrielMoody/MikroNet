package controller

import (
	"net/http"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/dto"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/pb"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type DashboardController interface {
	GetBusinessOwners(c *fiber.Ctx) error
	GetBusinessOwnerDetails(c *fiber.Ctx) error
	GetBlockedBusinessOwners(c *fiber.Ctx) error
	GetUnverifiedBusinessOwners(c *fiber.Ctx) error
	SetOwnerStatusVerified(c *fiber.Ctx) error
	SetDriverStatusVerified(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUserDetails(c *fiber.Ctx) error
	GetDrivers(c *fiber.Ctx) error
	GetDriverDetails(c *fiber.Ctx) error
	BlockAccount(c *fiber.Ctx) error
	UnblockAccount(c *fiber.Ctx) error
	GetReviews(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
}

type DashboardControllerImpl struct {
	DashboardService service.DashboardService
	PBDriver         pb.DriverServiceClient
	PBUser           pb.UserServiceClient
}

func (a *DashboardControllerImpl) SetDriverStatusVerified(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.PBDriver.SetStatusVerified(ctx, &pb.ReqByID{
		Id: id,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) SetOwnerStatusVerified(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.SetStatusVerified(ctx, id)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetUnverifiedBusinessOwners(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetUnverifiedBusinessOwners(ctx)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"count":  len(res),
			"owners": res,
		},
	})
}

func (a *DashboardControllerImpl) UnblockAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	accountId := c.Params("id")

	res, err := a.DashboardService.UnblockAccount(ctx, accountId)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetBusinessOwnerDetails(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.GetBusinessOwner(ctx, id)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetReviews(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.PBUser.GetReviews(ctx, &pb.Empty{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"reviews": res.Reviews,
			"count":   len(res.Reviews),
		},
	})
}

func (a *DashboardControllerImpl) GetReviewByID(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.PBUser.GetReviewsByID(ctx, &pb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
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

	res, err := a.DashboardService.BlockAccount(ctx, accountId)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetBlockedBusinessOwners(c *fiber.Ctx) error {
	ctx := c.Context()
	role := c.Query("role")

	res, err := a.DashboardService.GetBlockedBusinessOwners(ctx, role)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"count":    len(res),
			"accounts": res,
		},
	})
}

func (a *DashboardControllerImpl) GetBusinessOwners(c *fiber.Ctx) error {
	ctx := c.Context()

	res, err := a.DashboardService.GetBusinessOwners(ctx)

	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status": "error",
			"errors": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"count":  len(res),
			"owners": res,
		},
	})
}

func (a *DashboardControllerImpl) GetDriverDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := a.PBDriver.GetDriverDetails(c.Context(), &pb.ReqByID{
		Id: id,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"id":    res.Id,
			"name":  res.Name,
			"email": res.Email,
			"image": res.ImageUrl,
		},
	})
}

func (a *DashboardControllerImpl) GetDrivers(c *fiber.Ctx) error {
	var q dto.GetDriverQuery

	if err := c.QueryParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	res, err := a.PBDriver.GetDrivers(c.Context(), &pb.ReqDrivers{Verified: q.Verified})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"drivers": res.Drivers,
			"count":   len(res.Drivers),
		},
	})
}

func (a *DashboardControllerImpl) GetUsers(c *fiber.Ctx) error {
	res, err := a.PBUser.GetUsers(c.Context(), &pb.Empty{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data": fiber.Map{
			"users": res.Users,
			"count": len(res.Users),
		},
	})
}

func (a *DashboardControllerImpl) GetUserDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := a.PBUser.GetUserDetails(c.Context(), &pb.GetByIDRequest{
		Id: id,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func NewDashboardController(service service.DashboardService, driver pb.DriverServiceClient, user pb.UserServiceClient) DashboardController {
	return &DashboardControllerImpl{DashboardService: service, PBDriver: driver, PBUser: user}
}
