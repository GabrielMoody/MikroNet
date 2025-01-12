package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey;type:varchar(36)"`
	Email     string `gorm:"unique;type:varchar(255)"`
	Password  string
	Role      string `gorm:"type:enum('admin','user','driver','owner','government')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResetPassword struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"unique;type:varchar(36)"`
	User   User   `gorm:"foreignKey:UserID"`
	Code   string `gorm:"type:varchar(255)"`
}
