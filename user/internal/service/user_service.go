package service

import (
	"context"
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/helper"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"net/http"
	"time"
)

type UserService interface {
	GetUserDetails(c context.Context, id string) (res model.UserDetails, err *helper.ErrorStruct)
	EditUserDetails(c context.Context, id string, data dto.EditUserDetails) (res model.UserDetails, err *helper.ErrorStruct)
	DeleteUserDetails(c context.Context, id string) (err *helper.ErrorStruct)
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

func (a *userServiceImpl) EditUserDetails(c context.Context, id string, data dto.EditUserDetails) (res model.UserDetails, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	format := "02-01-2006"
	date, _ := time.Parse(format, data.DateOfBirth)

	resRepo, errRepo := a.repo.EditUserDetails(c, model.UserDetails{
		ID:          id,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		DateOfBirth: date,
		Gender:      data.Gender,
		Age:         int32(data.Age),
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) DeleteUserDetails(c context.Context, id string) (err *helper.ErrorStruct) {
	errRepo := a.repo.DeleteUserDetails(c, id)

	if errRepo != nil {
		return &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return nil
}

func (a *userServiceImpl) ReviewOrder(c context.Context, data dto.ReviewReq, userId string, driverId string) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, errRepo := a.repo.ReviewOrder(c, data, userId, driverId)

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
