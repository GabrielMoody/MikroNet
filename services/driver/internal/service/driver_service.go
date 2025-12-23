package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/dto"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/helper"
	"github.com/GabrielMoody/MikroNet/services/driver/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type DriverService interface {
	GetDriverDetails(c context.Context, id string) (res dto.GetDriverDetailsRes, err *helper.ErrorStruct)
	GetStatus(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
	SetStatus(c context.Context, id string, data dto.StatusReq) (res interface{}, err *helper.ErrorStruct)
	ConfirmOrder(c context.Context, order_id string, is_acepted bool) (res interface{}, err *helper.ErrorStruct)
}

type driverServiceImpl struct {
	repo repository.DriverRepo
	amqp *common.AMQP
}

func (a *driverServiceImpl) ConfirmOrder(c context.Context, order_id string, is_accepted bool) (res interface{}, err *helper.ErrorStruct) {
	data := dto.OrderConfirmation{
		IsAccepted: is_accepted,
	}

	b, _ := json.Marshal(data)

	a.amqp.PublishPersistent("order", "order.confirmation", b)

	return true, nil
}

func (a *driverServiceImpl) GetDriverDetails(c context.Context, id string) (res dto.GetDriverDetailsRes, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetDriverDetails(c, id)

	if errRepo != nil {
		var code int

		switch {
		case errors.Is(errRepo, helper.ErrNotFound):
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: code,
		}
	}

	return dto.GetDriverDetailsRes{
		ID:            resRepo.ID,
		Name:          resRepo.Name,
		LicenseNumber: resRepo.PlateNumber,
	}, nil
}

func (a *driverServiceImpl) GetStatus(c context.Context, id string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetStatus(c, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *driverServiceImpl) SetStatus(c context.Context, id string, data dto.StatusReq) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(&data); err != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, errRepo := a.repo.SetStatus(c, &data.IsOnline, id)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func NewDriverService(repo repository.DriverRepo, amqp *common.AMQP) DriverService {
	return &driverServiceImpl{
		repo: repo,
		amqp: amqp,
	}
}
