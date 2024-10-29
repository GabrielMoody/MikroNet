package service

import (
	"context"
	"github.com/GabrielMoody/mikroNet/profiles/internal/dto"
	"github.com/GabrielMoody/mikroNet/profiles/internal/helper"
	"github.com/GabrielMoody/mikroNet/profiles/internal/models"
	"github.com/GabrielMoody/mikroNet/profiles/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type ProfileService interface {
	GetUserService(c context.Context, id string) (res dto.UserRegistrationsResp, err *helper.ErrorStruct)
	EditUserService(c context.Context, id string, data dto.UserChangeProfileReq) (res string, err *helper.ErrorStruct)
	DeleteUserService(c context.Context, id string) (res string, err *helper.ErrorStruct)
	ChangePasswordService(c context.Context, id string, data dto.ChangePasswordReq) (res string, err *helper.ErrorStruct)
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
		ID:          resRepo.ID,
		FirstName:   resRepo.FirstName,
		LastName:    resRepo.LastName,
		Email:       resRepo.Email,
		PhoneNumber: resRepo.PhoneNumber,
		Role:        resRepo.Role,
		ImageUrl:    resRepo.ImageURL,
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
		ID:        id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Gender:    data.Gender,
		ImageURL:  data.Image,
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

func (a *ProfileServiceImpl) ChangePasswordService(c context.Context, id string, data dto.ChangePasswordReq) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := a.ProfileRepo.ChangePassword(c, data.OldPassword, data.NewPassword, id)

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
