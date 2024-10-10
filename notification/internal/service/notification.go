package service

import (
	"context"
	"github.com/GabrielMoody/mikroNet/notification/internal/dto"
	"github.com/GabrielMoody/mikroNet/notification/internal/model"
	"github.com/GabrielMoody/mikroNet/notification/internal/repository"
)

type NotificationService interface {
	Create(c context.Context, data dto.NotificationData) (string, error)
	Find(c context.Context, id string) (*model.Notification, error)
	FindAll(c context.Context) ([]*model.Notification, error)
}

type NotificationServiceImpl struct {
	repo repository.NotificationRepo
}

func (n NotificationServiceImpl) Create(c context.Context, data dto.NotificationData) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationServiceImpl) Find(c context.Context, id string) (*model.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationServiceImpl) FindAll(c context.Context) ([]*model.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func NewNotificationService(repo repository.NotificationRepo) NotificationService {
	return &NotificationServiceImpl{
		repo: repo,
	}
}
