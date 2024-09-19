package repository

import (
	"context"
	"github.com/GabrielMoody/mikroNet/driver/internal/model"
	"gorm.io/gorm"
	"time"
)

type DriverRepo interface {
	GetStatus(c context.Context, id string) (res interface{}, err error)
	SetStatus(c context.Context, status string, id string) (res interface{}, err error)
	GetRequest(c context.Context) (res interface{}, err error)
	AcceptRequest(c context.Context) (res interface{}, err error)
	GetAvailableSeats(c context.Context, id string) (res interface{}, err error)
	SetAvailableSeats(c context.Context, data model.Driver) (res interface{}, err error)
	GetTripHistories(c context.Context, id string) (res interface{}, err error)
}

type DriverRepoImpl struct {
	db *gorm.DB
}

func (a *DriverRepoImpl) GetTripHistories(c context.Context, id string) (res interface{}, err error) {
	row, err := a.db.WithContext(c).Table("trips").
		Select("trips.location, trips.destination, trips.trip_date, reviews.review, reviews.star").
		Joins("JOIN reviews ON trips.id = reviews.id").
		Where("trips.driver_id = ?", id).
		Rows()

	if err != nil {
		return nil, err
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
	var driver model.Driver

	if err := a.db.WithContext(c).Select("available_seats").Where("id = ?", id).Find(&driver).Error; err != nil {
		return nil, err
	}

	return driver.AvailableSeats, nil
}

func (a *DriverRepoImpl) SetAvailableSeats(c context.Context, data model.Driver) (res interface{}, err error) {
	if err := a.db.WithContext(c).Updates(data).Error; err != nil {
		return nil, err
	}

	return data.AvailableSeats, nil
}

func (a *DriverRepoImpl) GetStatus(c context.Context, id string) (res interface{}, err error) {
	var driver model.Driver

	if err := a.db.WithContext(c).Select("status").First(&driver, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return driver.Status, nil
}

func (a *DriverRepoImpl) SetStatus(c context.Context, status string, id string) (res interface{}, err error) {
	if err := a.db.WithContext(c).Where("id = ?", id).Updates(model.Driver{Status: status}).Error; err != nil {
		return nil, err
	}

	return status, nil
}

func (a *DriverRepoImpl) GetRequest(c context.Context) (res interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *DriverRepoImpl) AcceptRequest(c context.Context) (res interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func NewDriverRepo(db *gorm.DB) DriverRepo {
	return &DriverRepoImpl{
		db: db,
	}
}
