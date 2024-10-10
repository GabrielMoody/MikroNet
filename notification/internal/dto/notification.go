package dto

type NotificationData struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
	IsRead  bool   `json:"is_read"`
}
