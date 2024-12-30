package model

import "time"

type UserDetails struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	FirstName   string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Email       string    `gorm:"column:email" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	DateOfBirth time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	Age         int32     `gorm:"column:age" json:"age"`
	Gender      string    `gorm:"column:gender" json:"gender"`
}

type Review struct {
	ID       int    `gorm:"column:id;primaryKey" json:"id"`
	UserID   string `gorm:"column:user_id" json:"user_id"`
	DriverID string `gorm:"column:driver_id" json:"driver_id"`
	Comment  string `gorm:"column:comment" json:"comment"`
	Star     int    `gorm:"column:star" json:"star"`
}
