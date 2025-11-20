package dto

type (
	Hub struct {
		NotificationChannel map[string]chan NotificationData
	}
)
