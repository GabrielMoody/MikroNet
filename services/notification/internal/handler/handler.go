package handler

import (
	"github.com/GabrielMoody/mikroNet/notification/internal/dto"
	"github.com/GabrielMoody/mikroNet/notification/internal/sse"
	"github.com/gofiber/fiber/v2"
)

func NewHandler(r fiber.Router) {
	h := &dto.Hub{
		NotificationChannel: map[string]chan dto.NotificationData{},
	}

	//repo := repository.NewNotificationRepository(db)
	//service := service.NewNotificationService(repo)
	sseNotification := sse.NewNotificationSse(h)

	api := r.Group("/")

	api.Get("/sse/notification/:id", sseNotification.SseHandler)
	api.Post("/sse/notify/:id", sseNotification.NotifyUser)

}
