package controller

import (
	"github.com/GabrielMoody/mikroNet/dashboard/internal/middleware"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/pb"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/service"
	"github.com/gofiber/fiber/v2"
	"os"
)

type DashboardController interface {
	RegisterBusinessOwner(c *fiber.Ctx) error
	GetBusinessOwners(c *fiber.Ctx) error
	GetBusinessOwnerDetails(c *fiber.Ctx) error
	GetBlockedBusinessOwners(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUserDetails(c *fiber.Ctx) error
	GetDrivers(c *fiber.Ctx) error
	GetDriverDetails(c *fiber.Ctx) error
	BlockAccount(c *fiber.Ctx) error
	GetReviews(c *fiber.Ctx) error
	GetReviewByID(c *fiber.Ctx) error
}

type DashboardControllerImpl struct {
	DashboardService service.DashboardService
	PBDriver         pb.DriverServiceClient
	PBUser           pb.UserServiceClient
}

func (a *DashboardControllerImpl) GetBusinessOwnerDetails(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	res, err := a.DashboardService.GetBusinessOwner(ctx, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
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
			"status": "Error",
			"error":  err.Error(),
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
			"status": "Error",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) BlockAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	accountId := c.Params("accountId")

	token := c.Get("Authorization")
	payload, _ := middleware.GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))

	if payload["role"].(string) != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "Error",
			"error":  "Unauthorized",
		})
	}

	res, err := a.DashboardService.BlockAccount(ctx, accountId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetBlockedBusinessOwners(c *fiber.Ctx) error {
	ctx := c.Context()
	role := c.Query("role")

	res, err := a.DashboardService.GetBlockedBusinessOwners(ctx, role)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err,
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

	res, err := a.PBDriver.GetDriverDetails(c.Context(), &pb.ReqDriverDetails{
		Id: id,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) GetDrivers(c *fiber.Ctx) error {
	res, err := a.PBDriver.GetDrivers(c.Context(), &pb.Empty{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err.Error(),
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
			"status": "Error",
			"error":  err.Error(),
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Error",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "Success",
		"data":   res,
	})
}

func (a *DashboardControllerImpl) RegisterBusinessOwner(c *fiber.Ctx) error {
	//var req dto.OwnerRegistrationReq
	//ctx := c.Context()
	//
	//if err := c.BodyParser(&req); err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"status": "Error",
	//		"error":  err.Error(),
	//	})
	//}
	//
	//res, err := a.OwnerService.RegisterBusinessOwner(ctx, req)
	//
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"status": "Error",
	//		"error":  err.Err.Error(),
	//	})
	//}
	//
	//return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//	"status": "Success",
	//	"data":   res,
	//})
	return nil
}

func NewDashboardController(service service.DashboardService, driver pb.DriverServiceClient, user pb.UserServiceClient) DashboardController {
	return &DashboardControllerImpl{DashboardService: service, PBDriver: driver, PBUser: user}
}
