package rabbitmq

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/common"
)

func Init(url string) *common.AMQP {
	amqp := common.New(url)
	err := amqp.Connect(context.Background())

	if err != nil {
		panic(err)
	}

	return amqp
}
