package service

import (
	"context"
	"errors"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/helper"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/models"
	"github.com/GabrielMoody/mikroNet/dashboard/internal/repository"
	"net/http"
)

type DashboardService interface {
	GetBusinessOwners(c context.Context) (res []models.OwnerDetails, err *helper.ErrorStruct)
	GetBusinessOwner(c context.Context, id string) (res models.OwnerDetails, err *helper.ErrorStruct)
	GetBlockedBusinessOwners(c context.Context, role string) (res []models.OwnerDetails, err *helper.ErrorStruct)
	BlockAccount(c context.Context, accountId string) (res models.BlockedAccount, err *helper.ErrorStruct)
}

type DashboardServiceImpl struct {
	DashboardRepo repository.DashboardRepo
}

func (a *DashboardServiceImpl) GetBusinessOwner(c context.Context, id string) (res models.OwnerDetails, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetBusinessOwner(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) BlockAccount(c context.Context, accountId string) (res models.BlockedAccount, err *helper.ErrorStruct) {
	data := models.BlockedAccount{
		AccountID: accountId,
	}

	resRepo, errRepo := a.DashboardRepo.BlockAccount(c, data)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetBlockedBusinessOwners(c context.Context, role string) (res []models.OwnerDetails, err *helper.ErrorStruct) {
	if (role != "driver") && (role != "owner") && (role != "") {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errors.New("invalid role"),
		}
	}

	var resRepo []models.OwnerDetails
	var errRepo error

	if role == "" {
		resRepo, errRepo = a.DashboardRepo.GetAllBlockedAccount(c)
	} else {
		resRepo, errRepo = a.DashboardRepo.GetBlockedAccountRole(c, role)
	}

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *DashboardServiceImpl) GetBusinessOwners(c context.Context) (res []models.OwnerDetails, err *helper.ErrorStruct) {
	resRepo, errRepo := a.DashboardRepo.GetBusinessOwners(c)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func NewDashboardService(DashboardRepo repository.DashboardRepo) DashboardService {
	return &DashboardServiceImpl{
		DashboardRepo: DashboardRepo,
	}
}
