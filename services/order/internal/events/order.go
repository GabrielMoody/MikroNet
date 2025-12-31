package events

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/service"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rabbitmq/amqp091-go"
)

type handler func(context.Context, amqp091.Delivery) error

type Q struct {
	RoutingKey string
	Queue      string
	Handler    handler
}

type OrderEvents interface {
	Listen(c context.Context) error
}

type OrderEventsImpl struct {
	amqp_cons *common.AMQP
	service   service.OrderService
	queue     []Q
}

func (a *OrderEventsImpl) Listen(c context.Context) error {
	for _, qname := range a.queue {
		go func(queue Q) {
			messages, err := a.amqp_cons.Consume(queue.Queue, "order", queue.RoutingKey)

			if err != nil {
				log.Errorf("Error consuming queue: %s", err.Error())
			}

			for {
				select {
				case <-c.Done():
					log.Info("Stopping consumer:", qname.Queue)
					return
				case msg, ok := <-messages:
					if !ok {
						return
					}

					qname.Handler(c, msg)
				}
			}
		}(qname)

	}

	return nil
}

func NewEvents(service service.OrderService, amqp_cons *common.AMQP) OrderEvents {
	q := []Q{
		{
			RoutingKey: "order.created",
			Queue:      "order_created",
			Handler:    service.MakeOrder,
		},
		{
			RoutingKey: "order.notification",
			Queue:      "order_notifications",
		},
		{
			RoutingKey: "order.confirmation",
			Queue:      "order_confirmation",
			Handler:    service.ConfirmOrder,
		},
	}

	return &OrderEventsImpl{
		service:   service,
		amqp_cons: amqp_cons,
		queue:     q,
	}
}
