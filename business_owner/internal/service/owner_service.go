package service

import (
	"context"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/dto"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/helper"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/models"
	"github.com/GabrielMoody/mikroNet/business_owner/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type OwnerService interface {
	RegisterBusinessOwner(c context.Context, data dto.OwnerRegistrationReq) (res interface{}, err *helper.ErrorStruct)
	GetRatings(c context.Context, ownerId string) (res interface{}, err *helper.ErrorStruct)
	RegisterNewDriver(c context.Context, data dto.DriverRegistrationReq) (res interface{}, err *helper.ErrorStruct)
	GetDrivers(c context.Context, ownerId string) (res interface{}, err *helper.ErrorStruct)
	GetOwnerStatusVerified(c context.Context, ownerId string) (res interface{}, err *helper.ErrorStruct)
}

type OwnerServiceImpl struct {
	OwnerRepo repository.OwnerRepo
}

func (a *OwnerServiceImpl) RegisterBusinessOwner(c context.Context, data dto.OwnerRegistrationReq) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusInternalServerError,
		}
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)

	user := models.User{
		ID:          uuid.NewString(),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Password:    string(hash),
		Role:        "business_owner",
	}

	owner := models.Owner{
		ID:  user.ID,
		NIK: data.NIK,
	}

	resRepo, errRepo := a.OwnerRepo.RegisterBusinessOwner(c, user, owner)

	if errRepo != nil {
		return resRepo, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *OwnerServiceImpl) GetRatings(c context.Context, ownerId string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.OwnerRepo.GetRatings(c, ownerId)

	if errRepo != nil {
		return resRepo, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *OwnerServiceImpl) RegisterNewDriver(c context.Context, data dto.DriverRegistrationReq) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusInternalServerError,
		}
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)

	user := models.User{
		ID:          uuid.NewString(),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		PhoneNumber: data.PhoneNumber,
		Password:    string(hash),
		Role:        "driver",
	}

	driver := models.Driver{
		ID:                 user.ID,
		RegistrationNumber: data.RegistrationNumber,
	}

	resRepo, errRepo := a.OwnerRepo.RegisterNewDriver(c, user, driver)

	if errRepo != nil {
		return resRepo, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *OwnerServiceImpl) GetDrivers(c context.Context, ownerId string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.OwnerRepo.GetDrivers(c, ownerId)

	if errRepo != nil {
		return resRepo, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *OwnerServiceImpl) GetOwnerStatusVerified(c context.Context, ownerId string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.OwnerRepo.GetOwnerStatusVerified(c, ownerId)

	if errRepo != nil {
		return resRepo, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	if resRepo {
		return "Business owner is verified", nil
	}

	return "Business owner is not verified", nil
}

func NewOwnerService(OwnerRepo repository.OwnerRepo) OwnerService {
	return &OwnerServiceImpl{
		OwnerRepo: OwnerRepo,
	}
}
