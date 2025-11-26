package service

import (
	"context"
	"net/http"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/order/internal/helper"
	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/GabrielMoody/MikroNet/services/order/internal/repository"
)

type OrderService interface {
	MakeOrder(c context.Context, order_req dto.OrderReq) (res dto.OrderReq, err *helper.ErrorStruct)
}

type OrderServiceImpl struct {
	repo repository.OrderRepo
	amqp *common.AMQP
}

func (a *OrderServiceImpl) MakeOrder(c context.Context, order_req dto.OrderReq) (res dto.OrderReq, err *helper.ErrorStruct) {
	order := model.Order{
		PickupPoint: model.GeoPoint{
			Lat: res.PickupPoint.Lat,
			Lng: res.PickupPoint.Lng,
		},
		DropoffPoint: model.GeoPoint{
			Lat: res.DestPoint.Lat,
			Lng: res.DestPoint.Lng,
		},
		Status: "ORDER_CREATED",
	}

	_, errRepo := a.repo.MakeOrder(c, order)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return res, nil
}

func NewOrderService(repo repository.OrderRepo, amqp *common.AMQP) OrderService {
	return &OrderServiceImpl{
		repo: repo,
		amqp: amqp,
	}
}
