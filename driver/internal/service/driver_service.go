package service

import (
	"context"
	"github.com/GabrielMoody/mikroNet/driver/internal/dto"
	"github.com/GabrielMoody/mikroNet/driver/internal/helper"
	"github.com/GabrielMoody/mikroNet/driver/internal/model"
	"github.com/GabrielMoody/mikroNet/driver/internal/repository"
	"net/http"
)

type DriverService interface {
	GetStatus(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
	SetStatus(c context.Context, id string, data dto.StatusReq) (res interface{}, err *helper.ErrorStruct)
	GetRequest(c context.Context) (res interface{}, err *helper.ErrorStruct)
	AcceptRequest(c context.Context) (res interface{}, err *helper.ErrorStruct)
	GetAvailableSeats(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
	SetAvailableSeats(c context.Context, data dto.SeatReq, id string) (res interface{}, err *helper.ErrorStruct)
	GetTripHistories(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
}

type driverServiceImpl struct {
	repo repository.DriverRepo
}

func (a *driverServiceImpl) GetTripHistories(c context.Context, id string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetTripHistories(c, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *driverServiceImpl) GetAvailableSeats(c context.Context, id string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetAvailableSeats(c, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *driverServiceImpl) SetAvailableSeats(c context.Context, data dto.SeatReq, id string) (res interface{}, err *helper.ErrorStruct) {
	driver := model.Driver{
		ID:             id,
		AvailableSeats: data.Seat,
	}
	resRepo, errRepo := a.repo.SetAvailableSeats(c, driver)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
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
	resRepo, errRepo := a.repo.SetStatus(c, data.Status, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *driverServiceImpl) GetRequest(c context.Context) (res interface{}, err *helper.ErrorStruct) {
	//TODO implement me
	panic("implement me")
}

func (a *driverServiceImpl) AcceptRequest(c context.Context) (res interface{}, err *helper.ErrorStruct) {
	//TODO implement me
	panic("implement me")
}

func NewDriverService(repo repository.DriverRepo) DriverService {
	return &driverServiceImpl{
		repo: repo,
	}
}
