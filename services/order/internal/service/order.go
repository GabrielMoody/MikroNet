package service

import (
	"context"
	"log"
	"net/http"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/order/internal/helper"
	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/GabrielMoody/MikroNet/services/order/internal/repository"
	"github.com/redis/go-redis/v9"
)

type OrderService interface {
	MakeOrder(c context.Context, order_req dto.OrderReq) (res []redis.GeoLocation, err *helper.ErrorStruct)
}

type OrderServiceImpl struct {
	repo repository.OrderRepo
	amqp *common.AMQP
}

func (a *OrderServiceImpl) MakeOrder(c context.Context, order_req dto.OrderReq) (res []redis.GeoLocation, err *helper.ErrorStruct) {
	log.Println("Make an order from Order Service")
	order := model.Order{
		PickupPoint: model.GeoPoint{
			Lat: order_req.PickupPoint.Lat,
			Lng: order_req.PickupPoint.Lng,
		},
		DropoffPoint: model.GeoPoint{
			Lat: order_req.DestPoint.Lat,
			Lng: order_req.DestPoint.Lng,
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

	drivers, errRepo := a.repo.FindNearestDriver(c, order.PickupPoint)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return drivers, nil
}

func NewOrderService(repo repository.OrderRepo, amqp *common.AMQP) OrderService {
	return &OrderServiceImpl{
		repo: repo,
		amqp: amqp,
	}
}
