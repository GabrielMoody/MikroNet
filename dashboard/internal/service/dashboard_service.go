package service

import (
	"context"
	"errors"
	"os"

	"net/http"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/dto"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/helper"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type DashboardService interface {
	GetAllDrivers(c context.Context, q dto.GetDriverQuery) (res []models.Drivers, err error)
	GetAllPassengers(c context.Context) (res []models.Passengers, err error)
	GetDriverById(c context.Context, id string) (res models.Drivers, err error)
	GetPassengerById(c context.Context, id string) (res models.Passengers, err error)
	GetAllReviews(c context.Context) (res []models.Reviews, err error)
	GetAllBlockAccount(c context.Context) (res []models.BlockDriver, err error)
	GetReviewById(c context.Context, id string) (res models.Reviews, err error)
	GetAllHistories(c context.Context) (res []models.Histories, err error)
	EditAmountRoute(c context.Context, data dto.EditAmount, id string) (res models.Route, err error)
	BlockAccount(c context.Context, accountId string) (res models.BlockedAccount, err error)
	UnblockAccount(c context.Context, accountId string) (res string, err error)
	SetDriverStatusVerified(c context.Context, id string) (res string, err error)
	DeleteDriver(c context.Context, id string) (res string, err error)
	DeleteUser(c context.Context, id string) (res string, err error)
	AddRoute(c context.Context, data dto.AddRoute) (res models.Route, err error)
	MonthlyReport(c context.Context, query dto.MonthReport) (res dto.Report, err error)
	GetImage(c context.Context, id string) (res string, err error)
	GetRoutes(c context.Context) (res []models.Route, err error)
	DeleteRoute(c context.Context, id string) (res string, err error)
}

type DashboardServiceImpl struct {
	DashboardRepo repository.DashboardRepo
}

func (a *DashboardServiceImpl) DeleteRoute(c context.Context, id string) (res string, err error) {
	resRepo, errRepo := a.DashboardRepo.DeleteRoute(c, id)

	if errRepo != nil {
		var code int

		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetRoutes(c context.Context) (res []models.Route, err error) {
	resRepo, errRepo := a.DashboardRepo.GetRoutes(c)

	if errRepo != nil {
		var code int

		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetImage(c context.Context, id string) (res string, err error) {
	resRepo, errRepo := a.DashboardRepo.GetDriverByID(c, id)

	if errRepo != nil {
		var code int

		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	if resRepo.KTP == "" {
		return res, fiber.NewError(fiber.StatusNotFound, "Data KTP tidak ditemukan")
	}

	return resRepo.KTP, nil
}

func (a *DashboardServiceImpl) GetAllHistories(c context.Context) (res []models.Histories, err error) {
	resRepo, errRepo := a.DashboardRepo.GetAllTripHistories(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) EditAmountRoute(c context.Context, data dto.EditAmount, id string) (res models.Route, err error) {
	route := models.Route{
		Amount: data.Amount,
	}

	resRepo, errRepo := a.DashboardRepo.EditAmountRoute(c, route, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) MonthlyReport(c context.Context, query dto.MonthReport) (res dto.Report, err error) {
	resRepo, errRepo := a.DashboardRepo.MonthlyReport(c, query.Month)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) AddRoute(c context.Context, data dto.AddRoute) (res models.Route, err error) {
	resRepo, errRepo := a.DashboardRepo.AddRoute(c, models.Route{
		RouteName: data.RouteName,
		Amount:    data.Price,
	})

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) DeleteDriver(c context.Context, id string) (res string, err error) {
	resRepo, errRepo := a.DashboardRepo.DeleteDriver(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) DeleteUser(c context.Context, id string) (res string, err error) {
	resRepo, errRepo := a.DashboardRepo.DeleteUser(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) SetDriverStatusVerified(c context.Context, id string) (res string, err error) {
	resRepo, errRepo := a.DashboardRepo.SetDriverStatusVerified(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllBlockAccount(c context.Context) (res []models.BlockDriver, err error) {
	resRepo, errRepo := a.DashboardRepo.GetAllBlcokAccount(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllReviews(c context.Context) (res []models.Reviews, err error) {
	resRepo, errRepo := a.DashboardRepo.GetAllReview(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetReviewById(c context.Context, id string) (res models.Reviews, err error) {
	resRepo, errRepo := a.DashboardRepo.GetReviewById(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllDrivers(c context.Context, q dto.GetDriverQuery) (res []models.Drivers, err error) {
	resRepo, errRepo := a.DashboardRepo.GetAllDrivers(c, q.Verified)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	for i := range resRepo {
		if resRepo[i].ProfilePicture != "" {
			resRepo[i].ProfilePicture = os.Getenv("BASE_URL") + "/api/driver/images/" + resRepo[i].ID
		}

		if resRepo[i].KTP != "" {
			resRepo[i].KTP = os.Getenv("BASE_URL") + "/api/dashboard/ktp/" + resRepo[i].ID
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetAllPassengers(c context.Context) (res []models.Passengers, err error) {
	resRepo, errRepo := a.DashboardRepo.GetAllPassengers(c)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetDriverById(c context.Context, id string) (res models.Drivers, err error) {
	resRepo, errRepo := a.DashboardRepo.GetDriverByID(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	resRepo.ProfilePicture = os.Getenv("BASE_URL") + "/api/driver/images/" + resRepo.ID

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetPassengerById(c context.Context, id string) (res models.Passengers, err error) {
	resRepo, errRepo := a.DashboardRepo.GetPassengerByID(c, id)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) UnblockAccount(c context.Context, accountId string) (res string, err error) {
	resRepo, errRepo := a.DashboardRepo.UnblockAccount(c, accountId)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) BlockAccount(c context.Context, accountId string) (res models.BlockedAccount, err error) {
	data := models.BlockedAccount{
		UserID: accountId,
	}

	resRepo, errRepo := a.DashboardRepo.BlockAccount(c, data)

	if errRepo != nil {
		var code int
		switch {
		case errors.Is(errRepo, helper.ErrDuplicateEntry):
			code = http.StatusConflict
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, fiber.NewError(code, errRepo.Error())
	}

	return resRepo, nil
}

func NewDashboardService(DashboardRepo repository.DashboardRepo) DashboardService {
	return &DashboardServiceImpl{
		DashboardRepo: DashboardRepo,
	}
}
