package handler

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/order/internal/events"
	"github.com/GabrielMoody/MikroNet/services/order/internal/repository"
	"github.com/GabrielMoody/MikroNet/services/order/internal/service"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

func OrderHandler(db *gorm.DB, rdb *redis.Client, amqp *common.AMQP) {
	repo := repository.NewOrderRepo(db, rdb)
	serviceUser := service.NewOrderService(repo, amqp)
	events := events.NewUserController(serviceUser, amqp)

	events.Listen(context.Background())
}
