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
	ConfirmOrder(c context.Context, order model.Order) (model.Order, error)
	GetOrderByID(c context.Context, orderId int) (model.Order, error)
}

type OrderRepoImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (a *OrderRepoImpl) GetOrderByID(c context.Context, orderId int) (res model.Order, err error) {
	if err = a.db.WithContext(c).First(&res, orderId).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *OrderRepoImpl) ConfirmOrder(c context.Context, order model.Order) (model.Order, error) {
	if err := a.db.WithContext(c).Updates(order).Error; err != nil {
		return order, err
	}

	return order, nil
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
			Count:      5,
		},
		WithCoord: true,
		WithDist:  true,
	}

	drivers, err := a.rdb.GeoSearchLocation(c, "drivers:location", q).Result()

	if err != nil {
		return nil, err
	}

	// var availableDrivers []redis.GeoLocation
	// for _, d := range drivers {
	// 	isAvailable, _ := a.rdb.SIsMember(c, "drivers:available", d.Name).Result()
	// 	if isAvailable {
	// 		availableDrivers = append(availableDrivers, d)
	// 	}
	// }

	return drivers, nil
}

func NewOrderRepo(db *gorm.DB, rdb *redis.Client) OrderRepo {
	return &OrderRepoImpl{
		db:  db,
		rdb: rdb,
	}
}
