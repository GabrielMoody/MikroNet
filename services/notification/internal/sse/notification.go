package sse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/GabrielMoody/mikroNet/notification/internal/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

type NotificationSse struct {
	hub *dto.Hub
}

func NewNotificationSse(h *dto.Hub) *NotificationSse {
	return &NotificationSse{
		hub: h,
	}
}

func (n *NotificationSse) SseHandler(c *fiber.Ctx) error {
	userId := c.Params("id")
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Expose-Headers", "Content-Type")

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	n.hub.NotificationChannel[userId] = make(chan dto.NotificationData)

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprintf("id: %s\nevent: %s\n"+"data: \n\n", uuid.NewString(), "initial")
		_, _ = fmt.Fprint(w, event)
		_ = w.Flush()

		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				_, _ = fmt.Fprintf(w, "id: %s\ndata: ping \n\n", uuid.NewString())
				_ = w.Flush()
			case notification, ok := <-n.hub.NotificationChannel[userId]:
				if !ok {
					log.Println("notification channel closed")
					return
				}
				data, _ := json.Marshal(notification)
				event = fmt.Sprintf("id: %s\nevent: %s\n"+"data: %s\n\n", "notifcation updated", uuid.NewString(), data)

				_, _ = fmt.Fprint(w, event)
				_ = w.Flush()
			}
		}
	})

	return nil
}

func (n *NotificationSse) NotifyUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	message := c.Query("message", "Hello, this is a notification")

	if channel, ok := n.hub.NotificationChannel[userID]; ok {
		channel <- dto.NotificationData{
			ID:      userID,
			Title:   "Notification Test",
			Message: message,
			IsRead:  false,
		}
	}

	return c.SendString("Notification sent")
}
