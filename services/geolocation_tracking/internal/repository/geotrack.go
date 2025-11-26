package repository

import (
	"context"

	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/dto"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type GeoTrackRepository interface {
	SaveCurrentDriverLocation(c context.Context, location dto.Message) (model.DriverLocation, error)
}

type GeoTrackRepositoryImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (a *GeoTrackRepositoryImpl) SaveCurrentDriverLocation(c context.Context, location dto.Message) (res model.DriverLocation, err error) {
	a.rdb.GeoAdd(c, "drivers:location", &redis.GeoLocation{
		Name:      string(location.UserID),
		Longitude: location.Lng,
		Latitude:  location.Lat,
	})

	return res, nil
}

func NewGeoTrackRepository(db *gorm.DB, rdb *redis.Client) GeoTrackRepository {
	return &GeoTrackRepositoryImpl{
		db:  db,
		rdb: rdb,
	}
}
