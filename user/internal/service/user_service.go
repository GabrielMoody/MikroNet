package service

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/helper"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserService interface {
	GetRoutes(c context.Context) (res []model.Route, err *helper.ErrorStruct)
	OrderMikro(c context.Context, lat, lon, userId string) (res interface{}, err *helper.ErrorStruct)
	CarterMikro(c context.Context, route interface{}) (res interface{}, err *helper.ErrorStruct)
	GetTripHistories(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
	ReviewOrder(c context.Context, orderId string, data dto.ReviewReq) (res interface{}, err *helper.ErrorStruct)
}

type userServiceImpl struct {
	repo repository.UserRepo
}

func (a *userServiceImpl) ReviewOrder(c context.Context, orderId string, data dto.ReviewReq) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, errRepo := a.repo.ReviewOrder(c, data, orderId)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) GetRoutes(c context.Context) (res []model.Route, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetRoutes(c)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) OrderMikro(c context.Context, lat, lon, userId string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.OrderMikro(c, lat, lon, userId, nil)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) CarterMikro(c context.Context, route interface{}) (res interface{}, err *helper.ErrorStruct) {
	//TODO implement me
	panic("implement me")
}

func (a *userServiceImpl) GetTripHistories(c context.Context, id string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetTripHistories(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
