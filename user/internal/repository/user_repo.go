package repository

import (
	"context"
	"fmt"
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserRepo interface {
	CreateUser(c context.Context, user model.UserDetails) (model.UserDetails, error)
	GetUserDetails(c context.Context, id string) (model.UserDetails, error)
	GetAllUsers(c context.Context) ([]model.UserDetails, error)
	EditUserDetails(c context.Context, user model.UserDetails) (model.UserDetails, error)
	DeleteUserDetails(c context.Context, id string) error
	GetRoutes(c context.Context) ([]model.Route, error)
	OrderMikro(c context.Context, lat string, lon string, userId string, route interface{}) ([]dto.Orders, model.Order, error)
	CarterMikro(c context.Context, route interface{}) (interface{}, error)
	GetTripHistories(c context.Context, id string) (interface{}, error)
	ReviewOrder(c context.Context, data dto.ReviewReq, orderId string) (model.Review, error)
	findNearestDriver(c context.Context, lat string, lon string) ([]dto.Orders, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func (a *UserRepoImpl) GetAllUsers(c context.Context) (res []model.UserDetails, err error) {
	if err := a.db.WithContext(c).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (a *UserRepoImpl) GetUserDetails(c context.Context, id string) (res model.UserDetails, err error) {
	if err := a.db.WithContext(c).Find(&res, "id = ?", id).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *UserRepoImpl) EditUserDetails(c context.Context, user model.UserDetails) (model.UserDetails, error) {
	if err := a.db.WithContext(c).Updates(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (a *UserRepoImpl) DeleteUserDetails(c context.Context, id string) error {
	if err := a.db.WithContext(c).Delete(&model.UserDetails{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (a *UserRepoImpl) CreateUser(c context.Context, user model.UserDetails) (res model.UserDetails, err error) {
	if err = a.db.WithContext(c).Create(&user).Error; err != nil {
		return res, err
	}

	return user, nil
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

func (a *UserRepoImpl) OrderMikro(c context.Context, lat string, lon string, userId string, route interface{}) (driver []dto.Orders, res model.Order, err error) {
	drivers, err := a.findNearestDriver(c, lat, lon)

	if err != nil {
		return drivers, res, err
	}

	//order := model.Order{
	//	UserID: userId,
	//	//DriverID:        driver.DriverId,
	//	PickUpLocation:  "test",
	//	DropOffLocation: "test",
	//}
	//
	//if err := a.db.WithContext(c).Create(&order).Error; err != nil {
	//	return res, err
	//}

	return drivers, res, nil
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

func (a *UserRepoImpl) findNearestDriver(c context.Context, lat string, lon string) (res []dto.Orders, err error) {
	var d []dto.Orders

	err = a.db.WithContext(c).Table("driver_location").
		Select(fmt.Sprintf("drivers.id, users.first_name, users.last_name, drivers.registration_number, ST_Distance(\nCAST(FORMAT('SRID=4326;POINT(%%s %%s)', ST_X(location::geometry), ST_Y(location::geometry)) AS geography), \nCAST('SRID=4326;POINT(%s %s)' AS geography)) AS distance", lon, lat)).
		Joins("JOIN drivers ON drivers.id = driver_location.driver_id").
		Joins("JOIN users ON users.id = drivers.id").
		Where("drivers.status = ?", "on").
		Limit(5).Order("distance").Scan(&d).Error

	if err != nil {
		return res, err
	}
	var nd []dto.Orders
	for _, v := range d {
		f, _ := strconv.ParseFloat(v.Distance, 64)
		if f <= 1000 {
			nd = append(nd, v)
		}
	}

	return nd, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}
