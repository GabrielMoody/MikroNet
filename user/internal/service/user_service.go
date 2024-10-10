package service

import (
	"context"
	"fmt"
	"github.com/GabrielMoody/mikroNet/user/internal/dto"
	"github.com/GabrielMoody/mikroNet/user/internal/helper"
	"github.com/GabrielMoody/mikroNet/user/internal/model"
	"github.com/GabrielMoody/mikroNet/user/internal/repository"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type UserService interface {
	GetRoutes(c context.Context) (res []model.Route, err *helper.ErrorStruct)
	OrderMikro(c context.Context, lat, lon, userId string) (res interface{}, err *helper.ErrorStruct)
	CarterMikro(c context.Context, route interface{}) (res interface{}, err *helper.ErrorStruct)
	GetTripHistories(c context.Context, id string) (res interface{}, err *helper.ErrorStruct)
	ReviewOrder(c context.Context, orderId string, data dto.ReviewReq) (res interface{}, err *helper.ErrorStruct)
}

type userServiceImpl struct {
	repo repository.UserRepo
}

func (a *userServiceImpl) ReviewOrder(c context.Context, orderId string, data dto.ReviewReq) (res interface{}, err *helper.ErrorStruct) {
	if err := helper.Validate.Struct(data); err != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	resRepo, errRepo := a.repo.ReviewOrder(c, data, orderId)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: http.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) GetRoutes(c context.Context) (res []model.Route, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetRoutes(c)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func (a *userServiceImpl) OrderMikro(c context.Context, lat, lon, userId string) (res interface{}, err *helper.ErrorStruct) {
	drivers, _, errRepo := a.repo.OrderMikro(c, lat, lon, userId, nil)

	if errRepo != nil {
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}

	payload := map[string]string{
		"Message": "Ada orang yang pesan mikro",
		"Lat":     lat,
		"Lon":     lon,
	}

	for _, driver := range drivers {
		url := fmt.Sprintf("http://localhost:8015/sse/notify/%s?message=%s", driver.Id, payload)
		req, err := http.NewRequest("POST", url, nil)
		req.Header.Set("Content-Type", "text/event-stream")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Connection", "keep-alive")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return nil, &helper.ErrorStruct{
				Err:  err,
				Code: http.StatusInternalServerError,
			}
		}

		defer resp.Body.Close()

		// Log the response status
		log.Printf("Response for driver %s: %s", driver.Id, resp.Status)

		// Optional: Check if the response is successful
		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to notify driver %s: %s", driver.Id, resp.Status)
			continue // Handle as necessary
		}
	}

	return drivers, nil
}

func (a *userServiceImpl) CarterMikro(c context.Context, route interface{}) (res interface{}, err *helper.ErrorStruct) {
	//TODO implement me
	panic("implement me")
}

func (a *userServiceImpl) GetTripHistories(c context.Context, id string) (res interface{}, err *helper.ErrorStruct) {
	resRepo, errRepo := a.repo.GetTripHistories(c, id)

	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: http.StatusInternalServerError,
		}
	}

	return resRepo, nil
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}
