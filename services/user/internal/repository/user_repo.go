package repository

import (
	"context"

	"github.com/GabrielMoody/MikroNet/services/user/internal/helper"
	"github.com/GabrielMoody/MikroNet/services/user/internal/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	DeleteUser(c context.Context, id string) (model.User, error)
	GetUserDetails(c context.Context, id string) (model.User, error)
	MakeOrder(c context.Context, order model.Order) (model.Order, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func (a *UserRepoImpl) DeleteUser(c context.Context, id string) (res model.User, err error) {
	if err := a.db.WithContext(c).Delete(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *UserRepoImpl) GetUserDetails(c context.Context, id string) (res model.User, err error) {
	if err := a.db.WithContext(c).Table("users").
		Select("id, username, phone_number").
		Scan(&res).
		Error; err != nil {
		return res, helper.ErrNotFound
	}

	return res, nil
}

func (a *UserRepoImpl) MakeOrder(c context.Context, order model.Order) (res model.Order, err error) {
	if err := a.db.WithContext(c).Save(&order).Error; err != nil {
		return res, err
	}

	return res, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		db: db,
	}
}
