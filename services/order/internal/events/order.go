package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/order/internal/service"
)

type OrderEvents interface {
	Listen(c context.Context) error
}

type OrderEventsImpl struct {
	amqp_cons *common.AMQP
	service   service.OrderService
}

func (a *OrderEventsImpl) Listen(c context.Context) error {
	msgs, err := a.amqp_cons.Consume("order_created", "order", "order.created")

	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var order_req dto.OrderReq

			if err := json.Unmarshal(msg.Body, &order_req); err != nil {
				log.Fatal(err)
			}

			driver, _ := a.service.MakeOrder(c, order_req)

			log.Println(driver)

			msg.Ack(false)
		}
	}()

	select {}
}

func NewEvents(service service.OrderService, amqp_cons *common.AMQP) OrderEvents {
	return &OrderEventsImpl{
		service:   service,
		amqp_cons: amqp_cons,
	}
}
