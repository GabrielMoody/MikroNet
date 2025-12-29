package service

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/GabrielMoody/MikroNet/services/order/internal/repository"
	"github.com/rabbitmq/amqp091-go"
)

type OrderService interface {
	MakeOrder(c context.Context, msg amqp091.Delivery) error
	ConfirmOrder(c context.Context, msg amqp091.Delivery) error
	OrderNotification(c context.Context, msg amqp091.Delivery) error
	GetOrderByID(c context.Context, orderId int) (model.Order, error)
}

type OrderServiceImpl struct {
	repo     repository.OrderRepo
	amqp_pub *common.AMQP
}

func (a *OrderServiceImpl) GetOrderByID(c context.Context, orderId int) (res model.Order, err error) {
	res, errRepo := a.repo.GetOrderByID(c, orderId)

	if errRepo != nil {
		return res, err
	}

	return res, nil
}

func (a *OrderServiceImpl) MakeOrder(c context.Context, msg amqp091.Delivery) error {
	var order_req dto.OrderReq

	if errE := json.Unmarshal(msg.Body, &order_req); errE != nil {
		return errE
	}

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

	drivers, errRepo := a.repo.FindNearestDriver(c, order.PickupPoint)

	if errRepo != nil {
		return errRepo
	}
	driverId, _ := strconv.Atoi(drivers[0].Name)
	order.DriverID = int64(driverId)

	order, errRepo = a.repo.MakeOrder(c, order)

	if errRepo != nil {
		return errRepo
	}

	n := dto.OrderNotificationData{
		RecipientID: drivers[0].Name,
		OrderID:     int(order.ID),
		Title:       "New Order",
		PickupPoint: order_req.PickupPoint,
		DestPoint:   order_req.DestPoint,
	}

	b, _ := json.Marshal(n)

	errE := a.amqp_pub.PublishPersistent("order", "order.notification", append(b, '\n'))

	if errE != nil {
		log.Fatalf("failed to sent notification: %s", errE.Error())
	}

	return nil
}

func (a *OrderServiceImpl) ConfirmOrder(c context.Context, msg amqp091.Delivery) error {
	var req dto.OrderConfirmationReq

	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}

	var status string

	if req.IsConfirmed {
		status = "ORDER_ACCEPTED"
	} else {
		status = "ORDER_REJECTED"
	}

	order, err := a.repo.ConfirmOrder(c, model.Order{
		ID:     int64(req.OrderID),
		Status: status,
	})

	if err != nil {
		return err
	}

	n := dto.OrderNotificationData{
		RecipientID: strconv.Itoa(int(*order.UserID)),
		Title:       "Rejected",
	}

	b, _ := json.Marshal(n)

	errE := a.amqp_pub.PublishPersistent("order", "order.notification", append(b, '\n'))

	if errE != nil {
		log.Fatalf("failed to sent notification: %s", errE.Error())
	}

	return nil
}

func (a *OrderServiceImpl) OrderNotification(c context.Context, msg amqp091.Delivery) error {
	return nil
}

func NewOrderService(repo repository.OrderRepo, amqp_pub *common.AMQP) OrderService {
	return &OrderServiceImpl{
		repo:     repo,
		amqp_pub: amqp_pub,
	}
}
