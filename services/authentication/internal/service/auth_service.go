package service

import (
	"context"

	"github.com/GabrielMoody/mikronet-auth-service/internal/dto"
	"github.com/GabrielMoody/mikronet-auth-service/internal/helper"
	"github.com/GabrielMoody/mikronet-auth-service/internal/models"
	"github.com/GabrielMoody/mikronet-auth-service/internal/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateUserService(c context.Context, data dto.UserRegistrationsReq) (res int64, err *helper.ErrorStruct)
	CreateDriverService(c context.Context, data dto.DriverRegistrationsReq) (res int64, err *helper.ErrorStruct)
	LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct)
}

type AuthServiceImpl struct {
	AuthRepo repository.AuthRepo
}

func (a *AuthServiceImpl) CreateDriverService(c context.Context, data dto.DriverRegistrationsReq) (res int64, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return 0, &helper.ErrorStruct{
			Code:             fiber.StatusBadRequest,
			ValidationErrors: helper.ValidationError(errValidate),
		}
	}

	hashed, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return 0, &helper.ErrorStruct{
			Err:  errHash,
			Code: fiber.StatusInternalServerError,
		}
	}

	user := models.Authentication{
		Username: data.Email,
		Password: string(hashed),
		Role:     "driver",
		Driver: models.Driver{
			Name:        data.Name,
			PhoneNumber: data.PhoneNumber,
			PlateNumber: data.PlateNumber,
		},
	}

	resRepo, errRepo := a.AuthRepo.CreateDriver(c, user)

	if errRepo != nil {
		return res, helper.CheckError(errRepo)
	}

	return resRepo, nil
}

func (a *AuthServiceImpl) CreateUserService(c context.Context, data dto.UserRegistrationsReq) (res int64, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return 0, &helper.ErrorStruct{
			Code:             fiber.StatusBadRequest,
			ValidationErrors: helper.ValidationError(errValidate),
		}
	}

	hashed, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if errHash != nil {
		return 0, &helper.ErrorStruct{
			Err:  errHash,
			Code: fiber.StatusInternalServerError,
		}
	}

	user := models.Authentication{
		Username: data.Email,
		Password: string(hashed),
		Role:     "user",
		User: models.User{
			Fullname:    data.Name,
			PhoneNumber: data.PhoneNumber,
			Username:    data.Email,
		},
	}

	resRepo, errRepo := a.AuthRepo.CreateUser(c, user)

	if errRepo != nil {
		return res, helper.CheckError(errRepo)
	}

	return resRepo, nil
}

func (a *AuthServiceImpl) LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return res, &helper.ErrorStruct{
			Code:             fiber.StatusBadRequest,
			ValidationErrors: helper.ValidationError(errValidate),
		}
	}

	resRepo, errRepo := a.AuthRepo.LoginUser(c, data)

	if errRepo != nil {
		return res, helper.CheckError(errRepo)
	}

	return dto.UserRegistrationsResp{
		Email: resRepo.Username,
		Role:  resRepo.Role,
	}, nil
}

func NewAuthService(authRepo repository.AuthRepo) AuthService {
	return &AuthServiceImpl{
		AuthRepo: authRepo,
	}
}
