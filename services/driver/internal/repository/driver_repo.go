package repository

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/driver/internal/helper"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DriverRepo interface {
	GetDriverDetails(c context.Context, id string) (model.Driver, error)
	GetStatus(c context.Context, id string) (res interface{}, err error)
	SetStatus(c context.Context, status *bool, id string) (res interface{}, err error)
}

type DriverRepoImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func (a *DriverRepoImpl) CreateDriver(c context.Context, data model.Driver) (res model.Driver, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DriverRepoImpl) GetDriverDetails(c context.Context, id string) (res model.Driver, err error) {
	if err := a.db.WithContext(c).Table("driver_details").
		Select("driver_details.id as id, users.email, driver_details.name, driver_details.phone_number, driver_details.license_number, driver_details.sim, driver_details.verified, driver_details.profile_picture").
		Joins("JOIN users ON users.id = driver_details.id").
		Where("driver_details.id = ?", id).
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) GetStatus(c context.Context, id string) (res interface{}, err error) {
	var driver model.DriverStatus

	if err := a.db.WithContext(c).Select("is_online").First(&driver, "driver_id = ?", id).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return driver.IsOnline, nil
}

func (a *DriverRepoImpl) SetStatus(c context.Context, status *bool, id string) (res interface{}, err error) {
	if err := a.db.WithContext(c).Where("id = ?", id).Updates(model.DriverStatus{IsOnline: status}).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return status, nil
}

func NewDriverRepo(db *gorm.DB) DriverRepo {
	return &DriverRepoImpl{
		db: db,
	}
}
