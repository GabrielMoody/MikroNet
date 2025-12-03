package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GabrielMoody/MikroNet/services/common"
	"github.com/GabrielMoody/mikronet-user-service/internal/dto"
	"github.com/GabrielMoody/mikronet-user-service/internal/helper"
	"github.com/GabrielMoody/mikronet-user-service/internal/model"
	"github.com/GabrielMoody/mikronet-user-service/internal/repository"
)

type UserService interface {
	GetUserDetails(c context.Context, id string) (res model.User, err *helper.ErrorStruct)
	MakeOrder(c context.Context, order_req dto.OrderReq) (res dto.OrderReq, err *helper.ErrorStruct)
}

type userServiceImpl struct {
	repo repository.UserRepo
	amqp *common.AMQP
}

func (a *userServiceImpl) GetUserDetails(c context.Context, id string) (res model.User, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetUserDetails(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) MakeOrder(c context.Context, order_req dto.OrderReq) (res dto.OrderReq, err *helper.ErrorStruct) {
	// Published event order.created
	body, _ := json.Marshal(order_req)
	errEvent := a.amqp.PublishPersistent("order", "order.created", body)

	if errEvent != nil {
		return res, &helper.ErrorStruct{
			Err:  errEvent,
			Code: http.StatusInternalServerError,
		}
	}

	return res, nil
}

func NewUserService(repo repository.UserRepo, amqp *common.AMQP) UserService {
	return &userServiceImpl{
		repo: repo,
		amqp: amqp,
	}
}
