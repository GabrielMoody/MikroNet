package repository

import (
	"context"

	"github.com/GabrielMoody/mikronet-user-service/internal/helper"
	"github.com/GabrielMoody/mikronet-user-service/internal/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	DeleteUser(c context.Context, id string) (model.PassengerDetails, error)
	GetUserDetails(c context.Context, id string) (model.Users, error)
	ReviewOrder(c context.Context, data model.Review) (model.Review, error)
	Transaction(c context.Context, data model.Transaction) (model.Transaction, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func (a *UserRepoImpl) Transaction(c context.Context, data model.Transaction) (res model.Transaction, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *UserRepoImpl) DeleteUser(c context.Context, id string) (res model.PassengerDetails, err error) {
	if err := a.db.WithContext(c).Delete(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *UserRepoImpl) GetUserDetails(c context.Context, id string) (res model.Users, err error) {
	if err := a.db.WithContext(c).Table("users").
		Joins("JOIN passenger_details p ON p.id = users.id").
		Select("users.id, users.email, p.name").
		Scan(&res).
		Error; err != nil {
		return res, helper.ErrNotFound
	}

	return res, nil
}

func (a *UserRepoImpl) ReviewOrder(c context.Context, data model.Review) (res model.Review, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

//func (a *UserRepoImpl) findNearestDriver(c context.Context, lat string, lon string) (res []dto.Orders, err error) {
//	var d []dto.Orders
//
//	err = a.db.WithContext(c).Table("driver_location").
//		Select(fmt.Sprintf("drivers.id, users.first_name, users.last_name, drivers.registration_number, ST_Distance(\nCAST(FORMAT('SRID=4326;POINT(%%s %%s)', ST_X(location::geometry), ST_Y(location::geometry)) AS geography), \nCAST('SRID=4326;POINT(%s %s)' AS geography)) AS distance", lon, lat)).
//		Joins("JOIN drivers ON drivers.id = driver_location.driver_id").
//		Joins("JOIN users ON users.id = drivers.id").
//		Where("drivers.status = ?", "on").
//		Limit(5).Order("distance").Scan(&d).Error
//
//	if err != nil {
//		return res, err
//	}
//	var nd []dto.Orders
//	for _, v := range d {
//		f, _ := strconv.ParseFloat(v.Distance, 64)
//		if f <= 1000 {
//			nd = append(nd, v)
//		}
//	}
//
//	return nd, nil
//}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}
