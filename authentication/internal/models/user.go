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

type Driver struct {
	ID                 string    `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	OwnerID            string    `gorm:"column:owner_id" json:"owner_id"`
	RouteID            string    `gorm:"column:route_id" json:"route_id"`
	RegistrationNumber string    `gorm:"column:registration_number" json:"registration_number"`
	Status             string    `gorm:"column:status;default:off" json:"status"`
	AvailableSeats     int32     `gorm:"column:available_seats;default:9" json:"available_seats"`
	CreatedAt          time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ResetPassword struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"unique;size:191"`
	User   User   `gorm:"foreignKey:UserID"`
	Code   string
}
