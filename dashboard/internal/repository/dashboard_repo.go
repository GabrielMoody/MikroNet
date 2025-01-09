package repository

import (
	"context"
	"errors"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/helper"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/models"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	CreateBusinessOwner(c context.Context, data models.OwnerDetails) (models.OwnerDetails, error)
	GetBusinessOwners(c context.Context) ([]models.OwnerDetails, error)
	GetBusinessOwner(c context.Context, id string) (models.OwnerDetails, error)
	GetBlockedAccountRole(c context.Context, role string) ([]models.OwnerDetails, error)
	GetAllBlockedAccount(c context.Context) ([]models.OwnerDetails, error)
	GetUnverifiedBusinessOwners(c context.Context) ([]models.OwnerDetails, error)
	SetOwnerVerified(c context.Context, data models.OwnerDetails) (models.OwnerDetails, error)
	BlockAccount(c context.Context, data models.BlockedAccount) (models.BlockedAccount, error)
	UnblockAccount(c context.Context, id string) (string, error)
	IsBlocked(c context.Context, id string) (bool, error)
}

type DashboardRepoImpl struct {
	db *gorm.DB
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

func (a *DashboardRepoImpl) CreateBusinessOwner(c context.Context, data models.OwnerDetails) (res models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) SetOwnerVerified(c context.Context, data models.OwnerDetails) (res models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Updates(&data).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return data, nil
}

func (a *DashboardRepoImpl) GetUnverifiedBusinessOwners(c context.Context) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Find(&res, "verified = ?", false).Error; err != nil {
		return nil, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) UnblockAccount(c context.Context, id string) (res string, err error) {
	if err := a.db.WithContext(c).Delete(&models.BlockedAccount{}, "account_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, helper.ErrNotFound
		}
		return res, helper.ErrDatabase
	}

	return "Berhasil membuka blokir akun", nil
}

func (a *DashboardRepoImpl) GetBusinessOwner(c context.Context, id string) (res models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).First(&res, "id = ?", id).Error; err != nil {
		return res, helper.ErrNotFound
	}

	return res, nil
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

func (a *DashboardRepoImpl) GetBlockedAccountRole(c context.Context, role string) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Joins("JOIN blocked_accounts ON blocked_accounts.account_id = owner_details.id").Find(&res, "role = ?", role).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllBlockedAccount(c context.Context) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Joins("JOIN blocked_accounts ON blocked_accounts.account_id = owner_details.id").Find(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetBusinessOwners(c context.Context) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Find(&res).Error; err != nil {
		return res, helper.ErrDatabase
	}

	return res, nil
}

func NewDashboardRepo(db *gorm.DB) DashboardRepo {
	return &DashboardRepoImpl{
		db: db,
	}
}
