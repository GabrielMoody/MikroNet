package model

type Notification struct {
	ID      string `gorm:"column:id" json:"id"`
	UserID  string `gorm:"column:user_id" json:"user_id"`
	Title   string `gorm:"column:title" json:"title"`
	Message string `gorm:"column:message" json:"message"`
	IsRead  bool   `gorm:"column:is_read" json:"is_read"`
}
