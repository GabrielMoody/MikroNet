package service

import (
	"context"
	"encoding/json"
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
	repo     repository.OrderRepo
	amqp_pub *common.AMQP
}

func (a *OrderServiceImpl) MakeOrder(c context.Context, order_req dto.OrderReq) (res []redis.GeoLocation, err *helper.ErrorStruct) {
	log.Println("Make an order from Order Service")
	order := model.Order{
		UserID: &order_req.UserID,
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

	n := dto.OrderNotificationData{
		RecipientID: drivers[0].Name,
		Title:       "New Order",
		PickupPoint: order_req.PickupPoint,
		DestPoint:   order_req.DestPoint,
	}

	b, _ := json.Marshal(n)

	errE := a.amqp_pub.PublishPersistent("order", "order.notification", b)

	if errE != nil {
		log.Fatalf("failed to sent notification: %s", errE.Error())
	}

	return drivers, nil
}

func NewOrderService(repo repository.OrderRepo, amqp_pub *common.AMQP) OrderService {
	return &OrderServiceImpl{
		repo:     repo,
		amqp_pub: amqp_pub,
	}
}
