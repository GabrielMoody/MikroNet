package service

import (
	"context"
	"github.com/GabrielMoody/MikroNet/authentication/internal/dto"
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	"github.com/GabrielMoody/MikroNet/authentication/internal/models"
	"github.com/GabrielMoody/MikroNet/authentication/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProfileService interface {
	GetUserService(c context.Context, id string) (res dto.UserRegistrationsResp, err *helper.ErrorStruct)
	CreateUserService(c context.Context, data dto.UserRegistrationsReq) (res string, err *helper.ErrorStruct)
	LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct)
	EditUserService(c context.Context, id string, data dto.UserChangeProfileReq) (res string, err *helper.ErrorStruct)
	DeleteUserService(c context.Context, id string) (res string, err *helper.ErrorStruct)
	ChangePasswordService(c context.Context, oldPassword, newPassword, id string) (res string, err *helper.ErrorStruct)
}

type ProfileServiceImpl struct {
	ProfileRepo repository.ProfileRepo
}

func (a *ProfileServiceImpl) GetUserService(c context.Context, id string) (res dto.UserRegistrationsResp, err *helper.ErrorStruct) {
	resRepo, errRepo := a.ProfileRepo.GetUser(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return dto.UserRegistrationsResp{
		ID:           resRepo.ID,
		NamaLengkap:  resRepo.NamaLengkap,
		Email:        resRepo.Email,
		NomorTelepon: resRepo.NomorTelepon,
	}, nil
}

func (a *ProfileServiceImpl) CreateUserService(c context.Context, data dto.UserRegistrationsReq) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		return "", &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.ProfileRepo.CreateUser(c, models.User{
		ID:           uuid.NewString(),
		NamaLengkap:  data.NamaLengkap,
		Email:        data.Email,
		NomorTelepon: data.NomorTelepon,
		Password:     data.KataSandi,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *ProfileServiceImpl) LoginUserService(c context.Context, data dto.UserLoginReq) (res dto.UserRegistrationsResp, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.ProfileRepo.LoginUser(c, data)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return dto.UserRegistrationsResp{
		ID:           resRepo.ID,
		NamaLengkap:  resRepo.NamaLengkap,
		Email:        resRepo.Email,
		NomorTelepon: resRepo.NomorTelepon,
	}, nil
}

func (a *ProfileServiceImpl) EditUserService(c context.Context, id string, data dto.UserChangeProfileReq) (res string, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := a.ProfileRepo.UpdateUser(c, id, models.User{
		ID:           id,
		NamaLengkap:  data.NamaLengkap,
		Email:        data.Email,
		NomorTelepon: data.NomorTelepon,
		JenisKelamin: data.JenisKelamin,
	})

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *ProfileServiceImpl) DeleteUserService(c context.Context, id string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.ProfileRepo.DeleteUser(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *ProfileServiceImpl) ChangePasswordService(c context.Context, oldPassword, newPassword, id string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.ProfileRepo.ChangePassword(c, oldPassword, newPassword, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func NewProfileService(ProfileRepo repository.ProfileRepo) ProfileService {
	return &ProfileServiceImpl{
		ProfileRepo: ProfileRepo,
	}
}
