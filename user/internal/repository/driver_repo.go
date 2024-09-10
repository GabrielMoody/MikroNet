package repository

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"gorm.io/gorm"
)

type DriverRepo interface {
	GetStatus(c context.Context, id string) (res interface{}, err error)
	SetStatus(c context.Context, status string, id string) (res interface{}, err error)
	GetRequest(c context.Context) (res interface{}, err error)
	AcceptRequest(c context.Context) (res interface{}, err error)
}

type DriverRepoImpl struct {
	db *gorm.DB
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
