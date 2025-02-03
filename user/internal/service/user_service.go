package service

import (
	"context"
	"net/http"

	"github.com/GabrielMoody/mikronet-user-service/internal/dto"
	"github.com/GabrielMoody/mikronet-user-service/internal/helper"
	"github.com/GabrielMoody/mikronet-user-service/internal/model"
	"github.com/GabrielMoody/mikronet-user-service/internal/repository"
)

type UserService interface {
	GetUserDetails(c context.Context, id string) (res model.UserDetails, err *helper.ErrorStruct)
	ReviewOrder(c context.Context, data dto.ReviewReq, userId string, driverId string) (res interface{}, err *helper.ErrorStruct)
}

type userServiceImpl struct {
	repo repository.UserRepo
}

func (a *userServiceImpl) GetUserDetails(c context.Context, id string) (res model.UserDetails, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetUserDetails(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) ReviewOrder(c context.Context, data dto.ReviewReq, userId string, driverId string) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, errRepo := a.repo.ReviewOrder(c, model.Review{
		UserID:   userId,
		DriverID: driverId,
		Comment:  data.Comment,
		Star:     data.Star,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
