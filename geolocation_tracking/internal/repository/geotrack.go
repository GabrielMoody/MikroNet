package repository

import (
	"context"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/dto"
	"github.com/GabrielMoody/MikroNet/geolocation_tracking/internal/model"
	"gorm.io/gorm"
)

type GeoTrackRepository interface {
	SaveCurrentDriverLocation(c context.Context, location dto.Message) (model.DriverLocation, error)
}

type GeoTrackRepositoryImpl struct {
	db *gorm.DB
}

func (a *GeoTrackRepositoryImpl) SaveCurrentDriverLocation(c context.Context, location dto.Message) (res model.DriverLocation, err error) {
	if err := a.db.WithContext(c).Raw("INSERT INTO driver_locations (driver_id, location) VALUES (?, ST_SetSRID(ST_MakePoint(?, ?), 4326))"+
		" ON CONFLICT (driver_id) DO UPDATE SET location = EXCLUDED.location", location.UserID, location.Lng, location.Lat).Scan(res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func NewGeoTrackRepository(db *gorm.DB) GeoTrackRepository {
	return &GeoTrackRepositoryImpl{
		db: db,
	}
}
