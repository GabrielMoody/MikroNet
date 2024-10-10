package model

import "time"

type User struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	FirstName   string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Email       string    `gorm:"column:email" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	Password    string    `gorm:"column:password" json:"password"`
	DateOfBirth time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	Age         int32     `gorm:"column:age" json:"age"`
	Gender      string    `gorm:"column:gender" json:"gender"`
	Role        string    `gorm:"column:role" json:"role"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Driver struct {
	ID                 string `gorm:"column:id;primaryKey" json:"id"`
	OwnerID            string `gorm:"column:owner_id" json:"owner_id"`
	RegistrationNumber string `gorm:"column:registration_number" json:"registration_number"`
	Status             string `gorm:"column:status" json:"status"`
	Location           string `gorm:"column:location" json:"location"`
	AvailableSeats     int32  `gorm:"column:available_seats" json:"available_seats"`
}

type DriverLocation struct {
	DriverID  string    `gorm:"column:driver_id;primaryKey" json:"driver_id"`
	Location  string    `gorm:"column:location" json:"location"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type DriverLocationLogs struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	Location  string    `gorm:"column:location" json:"location"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}
