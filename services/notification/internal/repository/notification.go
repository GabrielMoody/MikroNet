package repository

import (
	"context"
	"github.com/GabrielMoody/mikroNet/notification/internal/dto"
	"github.com/GabrielMoody/mikroNet/notification/internal/model"
	"gorm.io/gorm"
)

type NotificationRepo interface {
	Create(c context.Context, data dto.NotificationData) (model.Notification, error)
	Find(c context.Context, id string) (model.Notification, error)
	FindAll(c context.Context) ([]model.Notification, error)
}

type NotificationImpl struct {
	db *gorm.DB
}

func (n *NotificationImpl) Create(c context.Context, data dto.NotificationData) (res model.Notification, err error) {
	notification := model.Notification{
		UserID:  data.ID,
		Title:   data.Title,
		Message: data.Message,
		IsRead:  data.IsRead,
	}

	if err := n.db.WithContext(c).Create(&notification).Error; err != nil {
		return model.Notification{}, err
	}

	return notification, nil
}

func (n *NotificationImpl) Find(c context.Context, id string) (res model.Notification, err error) {
	if err := n.db.WithContext(c).First(&res, "id = ?", id).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (n *NotificationImpl) FindAll(c context.Context) (res []model.Notification, err error) {
	if err := n.db.WithContext(c).Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func NewNotificationRepository(db *gorm.DB) NotificationRepo {
	return &NotificationImpl{
		db: db,
	}
}
