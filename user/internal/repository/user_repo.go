package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserRepo interface {
	GetRoutes(c context.Context) (res []model.Route, err error)
	OrderMikro(c context.Context, lat string, lon string, userId string, route interface{}) (res model.Order, err error)
	CarterMikro(c context.Context, route interface{}) (res interface{}, err error)
	GetTripHistories(c context.Context, id string) (res interface{}, err error)
	ReviewOrder(c context.Context, data dto.ReviewReq, orderId string) (res model.Review, err error)
	findNearestDriver(c context.Context, lat string, lon string) (res dto.Orders, err error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func (a *UserRepoImpl) ReviewOrder(c context.Context, data dto.ReviewReq, orderId string) (res model.Review, err error) {
	var order model.Order

	if err := a.db.WithContext(c).Where("order_id = ?", orderId).First(&order).Error; err != nil {
		return res, err
	}

	review := model.Review{
		UserID:   order.UserID,
		DriverID: order.DriverID,
		Review:   data.Review,
		Star:     data.Star,
	}

	if err := a.db.WithContext(c).Create(&review).Error; err != nil {
		return res, err
	}

	return review, nil
}

func (a *UserRepoImpl) GetRoutes(c context.Context) (res []model.Route, err error) {
	var routes []model.Route

	if err := a.db.WithContext(c).Find(&routes).Error; err != nil {
		return res, err
	}

	return routes, nil
}

func (a *UserRepoImpl) OrderMikro(c context.Context, lat string, lon string, userId string, route interface{}) (res model.Order, err error) {
	driver, err := a.findNearestDriver(c, lat, lon)

	if err != nil {
		return res, err
	}

	order := model.Order{
		UserID:          userId,
		DriverID:        driver.DriverId,
		PickUpLocation:  "test",
		DropOffLocation: "test",
	}

	if err := a.db.WithContext(c).Create(&order).Error; err != nil {
		return res, err
	}

	return order, nil
}

func (a *UserRepoImpl) CarterMikro(c context.Context, route interface{}) (res interface{}, err error) {
	panic("implement me")
}

func (a *UserRepoImpl) GetTripHistories(c context.Context, id string) (res interface{}, err error) {
	row, err := a.db.WithContext(c).Table("orders").
		Select("orders.location, orders.destination, orders.created_at, reviews.review, reviews.star").
		Joins("JOIN reviews ON orders.id = reviews.id").
		Where("orders.user_id = ?", id).
		Where("orders.status = completed").
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

func (a *UserRepoImpl) findNearestDriver(c context.Context, lat string, lon string) (res dto.Orders, err error) {
	var d dto.Orders

	err = a.db.WithContext(c).Table("driver_location").
		Select(fmt.Sprintf("drivers.id, users.first_name, users.last_name, drivers.registration_number, ST_Distance(\nCAST(FORMAT('SRID=4326;POINT(%s %s)', ST_Y(location::geometry), ST_X(location::geometry)) AS geography), \nCAST('SRID=4326;POINT(124.844728 1.493190)' AS geography)) AS distance", lon, lat)).
		Joins("JOIN drivers ON drivers.id = driver_location.driver_id").
		Joins("JOIN users ON users.id = drivers.id").
		Limit(5).Order("distance").Scan(&d).Error

	if err != nil {
		return res, err
	}

	distance, _ := strconv.ParseFloat(d.Distance, 5)

	if distance > 1000 {
		return res, errors.New("no nearest driver")
	}

	return d, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}
