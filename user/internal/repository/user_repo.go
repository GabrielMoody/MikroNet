package repository

import (
	"context"
	"fmt"
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"gorm.io/gorm"
	"strconv"
)

type UserRepo interface {
	CreateUser(c context.Context, user model.UserDetails) (model.UserDetails, error)
	GetUserDetails(c context.Context, id string) (model.UserDetails, error)
	GetAllUsers(c context.Context) ([]model.UserDetails, error)
	EditUserDetails(c context.Context, user model.UserDetails) (model.UserDetails, error)
	DeleteUserDetails(c context.Context, id string) error
	ReviewOrder(c context.Context, data dto.ReviewReq, userId string, driverId string) (model.Review, error)
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

func (a *UserRepoImpl) ReviewOrder(c context.Context, data dto.ReviewReq, userId string, driverId string) (res model.Review, err error) {
	review := model.Review{
		UserID:   userId,
		DriverID: driverId,
		Comment:  data.Comment,
		Star:     data.Star,
	}

	if err := a.db.WithContext(c).Create(&review).Error; err != nil {
		return res, err
	}

	return review, nil
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
