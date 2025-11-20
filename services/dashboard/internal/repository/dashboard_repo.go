package repository

import (
	"context"
	"errors"
	"time"

	"github.com/GabrielMoody/mikronet-dashboard-service/internal/dto"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/helper"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetAllDrivers(c context.Context, verified *bool) ([]models.Drivers, error)
	GetAllPassengers(c context.Context) ([]models.Passengers, error)
	GetDriverByID(c context.Context, id string) (models.Drivers, error)
	GetPassengerByID(c context.Context, id string) (models.Passengers, error)
	GetAllReview(c context.Context) ([]models.Reviews, error)
	GetReviewById(c context.Context, id string) (models.Reviews, error)
	GetAllTripHistories(c context.Context) ([]models.Histories, error)
	EditAmountRoute(c context.Context, data models.Route, id string) (models.Route, error)
	BlockAccount(c context.Context, data models.BlockedAccount) (models.BlockedAccount, error)
	UnblockAccount(c context.Context, id string) (string, error)
	IsBlocked(c context.Context, id string) (bool, error)
	GetAllBlcokAccount(c context.Context) ([]models.BlockDriver, error)
	SetDriverStatusVerified(c context.Context, id string) (string, error)
	DeleteDriver(c context.Context, id string) (string, error)
	DeleteUser(c context.Context, id string) (string, error)
	AddRoute(c context.Context, data models.Route) (models.Route, error)
	MonthlyReport(c context.Context, month int) (dto.Report, error)
	GetRoutes(c context.Context) ([]models.Route, error)
	DeleteRoute(c context.Context, id string) (string, error)
}

type DashboardRepoImpl struct {
	db *gorm.DB
}

func (a *DashboardRepoImpl) DeleteRoute(c context.Context, id string) (res string, err error) {
	if err := a.db.WithContext(c).Delete(&models.Route{}, "id = ?", id).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return "Berhasil menghapus rute!", nil
}

func (a *DashboardRepoImpl) GetRoutes(c context.Context) (res []models.Route, err error) {
	if err := a.db.WithContext(c).Find(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllTripHistories(c context.Context) (res []models.Histories, err error) {
	if err := a.db.WithContext(c).Table("transactions as t").
		Select("t.id as id, p.name as passenger_name, d.name as driver_name, t.amount as amount, r.route_name as route, t.created_at").
		Joins("JOIN passenger_details p on p.id = t.passenger_id").
		Joins("JOIN driver_details d on d.id = t.driver_id").
		Joins("JOIN routes r on r.id = d.route_id").
		Scan(&res).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) EditAmountRoute(c context.Context, data models.Route, id string) (res models.Route, err error) {
	if err := a.db.WithContext(c).Model(&res).Where("id = ?", id).Update("amount", data.Amount).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DashboardRepoImpl) MonthlyReport(c context.Context, month int) (res dto.Report, err error) {
	monthAgo := time.Now().AddDate(0, month*-1, 0)
	trips := []dto.RoutesReport{}
	common := dto.CommonReport{}

	sql := `
		WITH total_passenger AS (
    SELECT count(id) as total_passenger from passenger_details
    ), 
    total_driver AS (
        SELECT count(id) as total_driver from driver_details
    )
		SELECT COUNT(id) as total_trip, 
		SUM(amount) as total_revenue,
		(SELECT total_passenger FROM total_passenger) as total_passenger,
		(SELECT total_driver FROM total_driver) as total_driver FROM transactions;
	`

	if err := a.db.WithContext(c).Raw(sql).Scan(&common).Error; err != nil {
		return res, helper.ErrDatabase
	}

	if err := a.db.WithContext(c).Table("routes as r").
		Select("CONCAT('Rute ', r.id) as route, count(r.id) as total, sum(t.amount) as revenue").
		Joins("JOIN driver_details d on d.route_id = r.id").
		Joins("JOIN transactions t ON t.driver_id = d.id").
		Where("t.created_at >= ?", monthAgo).
		Group("r.id").
		Order("r.id").
		Scan(&trips).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return dto.Report{
		Common: common,
		Trips:  trips,
	}, nil
}

func (a *DashboardRepoImpl) AddRoute(c context.Context, data models.Route) (res models.Route, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DashboardRepoImpl) DeleteDriver(c context.Context, id string) (res string, err error) {
	q := a.db.WithContext(c).Delete(&models.DriverDetails{}, "id = ?", id)

	if q.Error != nil {
		return res, helper.ErrDatabase
	}

	if q.RowsAffected < 1 {
		return res, helper.ErrNotFound
	}

	return "Berhasil menghapus driver", nil
}

func (a *DashboardRepoImpl) DeleteUser(c context.Context, id string) (res string, err error) {
	q := a.db.WithContext(c).Delete(&models.PassengerDetails{}, "id = ?", id)

	if q.Error != nil {
		return res, helper.ErrDatabase
	}

	if q.RowsAffected < 1 {
		return res, helper.ErrNotFound
	}

	return "Berhasil menghapus passenger", nil
}

func (a *DashboardRepoImpl) SetDriverStatusVerified(c context.Context, id string) (res string, err error) {
	if err := a.db.WithContext(c).Model(&models.DriverDetails{}).Where("id = ?", id).Update("verified", true).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return "Berhasil memverifikasi driver", nil
}

func (a *DashboardRepoImpl) GetAllReview(c context.Context) (res []models.Reviews, err error) {
	if err := a.db.WithContext(c).Table("reviews").
		Select("reviews.id, p.name AS passenger_name, d.name AS driver_name, reviews.comment AS comment, reviews.star AS star").
		Joins("JOIN passenger_details p ON reviews.passenger_id = p.id").
		Joins("JOIN driver_details d ON reviews.driver_id = d.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetReviewById(c context.Context, id string) (res models.Reviews, err error) {
	if err := a.db.WithContext(c).Table("reviews").
		Select("reviews.id, p.name AS passenger_name, d.name AS driver_name, reviews.comment AS comment, reviews.star AS star").
		Joins("JOIN passenger_details p ON reviews.passenger_id = p.id").
		Joins("JOIN driver_details d ON reviews.driver_id = d.id").
		Where("reviews.id = ?", id).
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllDrivers(c context.Context, verified *bool) (res []models.Drivers, err error) {
	if verified == nil {
		if err := a.db.WithContext(c).Table("driver_details as d").
			Select("d.id as id, u.email, d.name, d.phone_number, d.license_number, d.sim, d.verified, d.profile_picture, d.ktp, d.status as status").
			Joins("JOIN users u ON u.id = d.id").
			Scan(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	} else {
		if err := a.db.WithContext(c).Table("driver_details").
			Select("driver_details.id as id, users.email, driver_details.name, driver_details.phone_number, driver_details.license_number, driver_details.sim, driver_details.verified, driver_details.profile_picture, driver_details.status as status").
			Joins("JOIN users ON users.id = driver_details.id").
			Where("verified = ?", verified).
			Scan(&res).Error; err != nil {
			return res, helper.ErrDatabase
		}
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllPassengers(c context.Context) (res []models.Passengers, err error) {
	if err := a.db.WithContext(c).Table("passenger_details").
		Select("passenger_details.id as id, users.email, passenger_details.name, passenger_details.date_of_birth, passenger_details.age").
		Joins("JOIN users ON users.id = passenger_details.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetDriverByID(c context.Context, id string) (res models.Drivers, err error) {
	if err := a.db.WithContext(c).Table("driver_details as d").
		Select("d.id as id, u.email, d.name, d.phone_number, d.license_number, d.sim, d.verified, d.profile_picture, d.ktp").
		Joins("JOIN users u ON u.id = d.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetPassengerByID(c context.Context, id string) (res models.Passengers, err error) {
	if err := a.db.WithContext(c).Table("passenger_details").
		Select("passenger_details.id as id, users.email, passenger_details.name").
		Joins("JOIN users ON users.id = passenger_details.id").
		Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllBlcokAccount(c context.Context) (res []models.BlockDriver, err error) {
	if err := a.db.WithContext(c).Table("blocked_accounts").
		Select("user_id as id, users.email as email, driver_details.name as name").
		Joins("JOIN users ON users.id = user_id").
		Joins("JOIN driver_details ON driver_details.id = user_id").Scan(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) IsBlocked(c context.Context, id string) (bool, error) {
	var res models.BlockedAccount
	if err := a.db.WithContext(c).First(&res, "account_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, helper.ErrDatabase
	}

	return true, nil
}

func (a *DashboardRepoImpl) UnblockAccount(c context.Context, id string) (res string, err error) {
	q := a.db.WithContext(c).Delete(&models.BlockedAccount{}, "user_id = ?", id)

	if q.Error != nil {
		return res, helper.ErrDatabase
	}

	if q.RowsAffected < 1 {
		return res, helper.ErrNotFound
	}

	return "Berhasil membuka blokir akun", nil
}

func (a *DashboardRepoImpl) BlockAccount(c context.Context, data models.BlockedAccount) (res models.BlockedAccount, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return res, helper.ErrDuplicateEntry
		}
		return res, helper.ErrDatabase
	}

	return res, nil
}

func NewDashboardRepo(db *gorm.DB) DashboardRepo {
	return &DashboardRepoImpl{
		db: db,
	}
}
