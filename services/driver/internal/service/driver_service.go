package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/GabrielMoody/mikronet-driver-service/internal/dto"
	"github.com/GabrielMoody/mikronet-driver-service/internal/helper"
	"github.com/GabrielMoody/mikronet-driver-service/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type DriverService interface {
	GetDriverDetails(c context.Context, id string) (res dto.GetDriverDetailsRes, err *helper.ErrorStruct)
	GetStatus(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
	SetStatus(c context.Context, id string, data dto.StatusReq) (res interface{}, err *helper.ErrorStruct)
}

type driverServiceImpl struct {
	repo repository.DriverRepo
}

func (a *driverServiceImpl) GetDriverDetails(c context.Context, id string) (res dto.GetDriverDetailsRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetDriverDetails(c, id)

	if errRepo != nil {
		var code int

		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return dto.GetDriverDetailsRes{
		ID:             resRepo.ID,
		Name:           resRepo.Name,
		Email:          resRepo.Email,
		LicenseNumber:  resRepo.LicenseNumber,
		SIM:            resRepo.SIM,
		ProfilePicture: resRepo.ProfilePicture,
	}, nil
}

func (a *driverServiceImpl) GetStatus(c context.Context, id string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetStatus(c, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *driverServiceImpl) SetStatus(c context.Context, id string, data dto.StatusReq) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(&data); err != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, errRepo := a.repo.SetStatus(c, data.Status, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func NewDriverService(repo repository.DriverRepo) DriverService {
	return &driverServiceImpl{
		repo: repo,
	}
}
