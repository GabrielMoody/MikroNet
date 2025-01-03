package repository

import (
	"context"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/models"
	"gorm.io/gorm"
)

type DashboardRepo interface {
	GetBusinessOwners(c context.Context) ([]models.OwnerDetails, error)
	GetBusinessOwner(c context.Context, id string) (models.OwnerDetails, error)
	GetBlockedAccountRole(c context.Context, role string) ([]models.OwnerDetails, error)
	GetAllBlockedAccount(c context.Context) ([]models.OwnerDetails, error)
	BlockAccount(c context.Context, data models.BlockedAccount) (models.BlockedAccount, error)
}

type DashboardRepoImpl struct {
	db *gorm.DB
}

func (a *DashboardRepoImpl) GetBusinessOwner(c context.Context, id string) (res models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Find(&res, "id = ?", id).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *DashboardRepoImpl) BlockAccount(c context.Context, data models.BlockedAccount) (res models.BlockedAccount, err error) {
	if err := a.db.WithContext(c).Create(&data).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetBlockedAccountRole(c context.Context, role string) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Joins("JOIN blocked_accounts ON blocked_accounts.account_id = owner_details.id").Find(&res, "role = ?", role).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetAllBlockedAccount(c context.Context) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Joins("JOIN blocked_accounts ON blocked_accounts.account_id = owner_details.id").Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (a *DashboardRepoImpl) GetBusinessOwners(c context.Context) (res []models.OwnerDetails, err error) {
	if err := a.db.WithContext(c).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func NewDashboardRepo(db *gorm.DB) DashboardRepo {
	return &DashboardRepoImpl{
		db: db,
	}
}
