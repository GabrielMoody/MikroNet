package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/GabrielMoody/mikronet-driver-service/internal/helper"
	"github.com/GabrielMoody/mikronet-driver-service/internal/model"
	"github.com/GabrielMoody/mikronet-driver-service/internal/pb"
	"gorm.io/gorm"
)

type DriverRepo interface {
	CreateDriver(c context.Context, data model.DriverDetails) (model.DriverDetails, error)
	GetAllDrivers(c context.Context, verified *pb.ReqDrivers) ([]model.DriverDetails, error)
	GetDriverDetails(c context.Context, id string) (model.Drivers, error)
	EditDriverDetails(c context.Context, user model.DriverDetails) (model.DriverDetails, error)
	DeleteDriver(c context.Context, id string) (model.DriverDetails, error)
	GetStatus(c context.Context, id string) (res interface{}, err error)
	SetStatus(c context.Context, status string, id string) (res interface{}, err error)
	GetTripHistories(c context.Context, id string) (res []model.Histories, err error)
	GetAllDriverLastSeen(c context.Context) (res []model.DriverDetails, err error)
	SetLastSeen(c context.Context, id string) (res *time.Time, err error)
	GetQrisData(c context.Context, id string) (res *string, err error)
}

type DriverRepoImpl struct {
	db *gorm.DB
}

func generateQrisData(id string) *string {
	qr := fmt.Sprintf("0002010102115802ID6006Manado6208%s530336054060006304A1B2", id)
	return &qr
}

func (a *DriverRepoImpl) GetQrisData(c context.Context, id string) (res *string, err error) {
	if err := a.db.WithContext(c).Model(&model.DriverDetails{}).Select("qris_data").Where("id = ?", id).Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	if res == nil {
		res = generateQrisData(id)
		if err := a.db.WithContext(c).Model(&model.DriverDetails{}).Where("id = ?", id).Update("qris_data", res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	}

	return res, nil
}

func (a *DriverRepoImpl) GetAllDriverLastSeen(c context.Context) (res []model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Find(&res).Where("last_seen >= ?", time.Now().Add(-5*time.Minute)).Scan(&res).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) SetLastSeen(c context.Context, id string) (res *time.Time, err error) {
	if err := a.db.WithContext(c).Model(&model.DriverDetails{}).Where("id = ?", id).Update("last_seen", time.Now()).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) DeleteDriver(c context.Context, id string) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Delete(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) CreateDriver(c context.Context, data model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DriverRepoImpl) GetAllDrivers(c context.Context, verified *pb.ReqDrivers) (res []model.DriverDetails, err error) {
	switch v := verified.Verified.(type) {
	case *pb.ReqDrivers_IsVerified:
		if err := a.db.WithContext(c).Find(&res, "verified = ?", v.IsVerified).Error; err != nil {
			return res, helper.ErrDatabase
		}
	case *pb.ReqDrivers_NotVerified:
		if err := a.db.WithContext(c).Find(&res, "verified = ?", v.NotVerified).Error; err != nil {
			return res, helper.ErrDatabase
		}
	default:
		if err := a.db.WithContext(c).Find(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	}

	return res, nil
}

func (a *DriverRepoImpl) GetDriverDetails(c context.Context, id string) (res model.Drivers, err error) {
	if err := a.db.WithContext(c).Table("driver_details").
		Select("driver_details.id as id, users.email, driver_details.name, driver_details.phone_number, driver_details.license_number, driver_details.sim, driver_details.verified, driver_details.profile_picture").
		Joins("JOIN users ON users.id = driver_details.id").
		Where("driver_details.id = ?", id).
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DriverRepoImpl) EditDriverDetails(c context.Context, driver model.DriverDetails) (res model.DriverDetails, err error) {
	if err := a.db.WithContext(c).Updates(&driver).Error; err != nil {
		return driver, helper.ErrDatabase
	}

	return driver, nil
}

func (a *DriverRepoImpl) GetTripHistories(c context.Context, id string) (res []model.Histories, err error) {
	if err := a.db.WithContext(c).Table("transactions").
		Select("transactions.id as id, passenger_details.name as passenger_name, driver_details.name as driver_name, transactions.amount as amount, transactions.created_at").
		Joins("JOIN passenger_details on passenger_details.id = transactions.passenger_id").
		Joins("JOIN driver_details on driver_details.id = transactions.driver_id").
		Where("driver_id = ?", id).Scan(&res).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return res, nil
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
