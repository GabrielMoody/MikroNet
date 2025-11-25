package service

import (
	"context"
	"net/http"

	"github.com/GabrielMoody/mikronet-user-service/internal/dto"
	"github.com/GabrielMoody/mikronet-user-service/internal/helper"
	"github.com/GabrielMoody/mikronet-user-service/internal/model"
	"github.com/GabrielMoody/mikronet-user-service/internal/repository"
)

type UserService interface {
	GetUserDetails(c context.Context, id string) (res model.User, err *helper.ErrorStruct)
	MakeOrder(c context.Context, order_req dto.OrderReq) (res dto.OrderReq, err *helper.ErrorStruct)
}

type userServiceImpl struct {
	repo repository.UserRepo
	amqp *common.AMQP
}

func (a *userServiceImpl) GetUserDetails(c context.Context, id string) (res model.User, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetUserDetails(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) MakeOrder(c context.Context, order_req dto.OrderReq) (res dto.OrderReq, err *helper.ErrorStruct) {
	order := model.Order{
		UserID: order_req.UserID,
		PickupPoint: model.GeoPoint{
			Lat: order_req.PickupPoint.Lat,
			Lng: order_req.PickupPoint.Lng,
		},
		DropoffPoint: model.GeoPoint{
			Lat: order_req.DestPoint.Lat,
			Lng: order_req.DestPoint.Lng,
		},
		Status: "ORDER_PENDING",
	}
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
