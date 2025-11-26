package rabbitmq

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/common"
)

func Init(url string) *common.AMQP {
	amqp := common.New("amqp://admin:admin123@rabbitmq:5672/")
	err := amqp.Connect(context.Background())

	if err != nil {
		panic(err)
	}

	// Declare order exchange
	amqp.DeclareExchange("order", "topic")

	return amqp
}
