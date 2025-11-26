package repository

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/order/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type OrderRepo interface {
	MakeOrder(c context.Context, order model.Order) (model.Order, error)
	FindNearestDriver(c context.Context, pickup_point model.GeoPoint) ([]redis.GeoLocation, error)
}

type OrderRepoImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (a *OrderRepoImpl) MakeOrder(c context.Context, order model.Order) (res model.Order, err error) {
	if err := a.db.WithContext(c).Save(&order).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *OrderRepoImpl) FindNearestDriver(c context.Context, pickup_point model.GeoPoint) ([]redis.GeoLocation, error) {
	q := &redis.GeoSearchLocationQuery{
		GeoSearchQuery: redis.GeoSearchQuery{
			Latitude:   pickup_point.Lat,
			Longitude:  pickup_point.Lng,
			Radius:     50,
			RadiusUnit: "km",
			Sort:       "ASC",
			Count:      1,
		},
		WithCoord: true,
		WithDist:  true,
	}

	drivers, err := a.rdb.GeoSearchLocation(c, "driver:location", q).Result()

	if err != nil {
		return nil, err
	}

	return drivers, nil
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &OrderRepoImpl{
		db: db,
	}
}
