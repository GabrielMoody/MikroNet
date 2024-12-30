package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResetPassword struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"unique;size:191"`
	User   User   `gorm:"foreignKey:UserID"`
	Code   string
}
