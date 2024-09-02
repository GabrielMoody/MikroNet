package models

import (
	"time"
)

type User struct {
	ID          string `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string `gorm:"unique"`
	Gender      string
	DateOfBirth *time.Time `gorm:"type:date"`
	Age         int
	Role        string `gorm:"default:'user'"`
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ResetPassword struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"unique;size:191"`
	User   User   `gorm:"foreignKey:UserID"`
	Code   string
}
