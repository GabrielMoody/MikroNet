package repository

import (
	"context"
	"github.com/GabrielMoody/mikroNet/driver/internal/helper"
	"github.com/GabrielMoody/mikroNet/driver/internal/model"
	"gorm.io/gorm"
	"time"
)

type DriverRepo interface {
	CreateDriver(c context.Context, data model.DriverDetails) (model.DriverDetails, error)
	GetAllDrivers(c context.Context, verified *bool) ([]model.DriverDetails, error)
	GetDriverDetails(c context.Context, id string) (model.DriverDetails, error)
	EditDriverDetails(c context.Context, user model.DriverDetails) (model.DriverDetails, error)
	GetStatus(c context.Context, id string) (res interface{}, err error)
	SetStatus(c context.Context, status string, id string) (res interface{}, err error)
	SetVerified(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error)
	GetAvailableSeats(c context.Context, id string) (res interface{}, err error)
	SetAvailableSeats(c context.Context, data model.DriverDetails) (res interface{}, err error)
	GetTripHistories(c context.Context, id string) (res interface{}, err error)
}

type DriverRepoImpl struct {
	db *gorm.DB
}

func (a *DriverRepoImpl) SetVerified(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Updates(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DriverRepoImpl) CreateDriver(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DriverRepoImpl) GetAllDrivers(c context.Context, verified *bool) (res []model.DriverDetails, err error) {
	if verified != nil {
		if err := a.db.WithContext(c).Find(&res, "verified = ?", verified).Error; err != nil {
			return res, helper.ErrDatabase
		}
	} else {
		if err := a.db.WithContext(c).Find(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	}

	return res, nil
}

func (a *DriverRepoImpl) GetDriverDetails(c context.Context, id string) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).First(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrNotFound
	}

	return res, nil
}

func (a *DriverRepoImpl) EditDriverDetails(c context.Context, user model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Updates(&user).Error; err != nil {
		return user, helper.ErrDatabase
	}

	return user, nil
}

func (a *DriverRepoImpl) GetTripHistories(c context.Context, id string) (res interface{}, err error) {
	row, err := a.db.WithContext(c).Table("trips").
		Select("trips.location, trips.destination, trips.trip_date, reviews.review, reviews.star").
		Joins("JOIN reviews ON trips.id = reviews.id").
		Where("trips.driver_id = ?", id).
		Rows()

	if err != nil {
		return nil, helper.ErrDatabase
	}

	defer row.Close()

	type data struct {
		Location    string
		Destination string
		TripDate    time.Time
		Review      string
		Star        int64
	}

	var d data
	var trip []data

	for row.Next() {
		_ = row.Scan(&d.Location, &d.Destination, &d.TripDate, &d.Review, &d.Star)
		trip = append(trip, d)
	}

	return trip, nil
}

func (a *DriverRepoImpl) GetAvailableSeats(c context.Context, id string) (res interface{}, err error) {
	var driver model.DriverDetails

	if err := a.db.WithContext(c).Select("available_seats").Where("id = ?", id).Find(&driver).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return driver.AvailableSeats, nil
}

func (a *DriverRepoImpl) SetAvailableSeats(c context.Context, data model.DriverDetails) (res interface{}, err error) {
	if err := a.db.WithContext(c).Updates(&data).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return data.AvailableSeats, nil
}

func (a *DriverRepoImpl) GetStatus(c context.Context, id string) (res interface{}, err error) {
	var driver model.DriverDetails

	if err := a.db.WithContext(c).Select("status").First(&driver, "id = ?", id).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return driver.Status, nil
}

func (a *DriverRepoImpl) SetStatus(c context.Context, status string, id string) (res interface{}, err error) {
	if err := a.db.WithContext(c).Where("id = ?", id).Updates(model.DriverDetails{Status: status}).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return status, nil
}

func NewDriverRepo(db *gorm.DB) DriverRepo {
	return &DriverRepoImpl{
		db: db,
	}
}
