package models

import "time"

type User struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	FirstName   string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Email       string    `gorm:"column:email;not null" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number;not null" json:"phone_number"`
	Password    string    `gorm:"column:password;not null" json:"password"`
	DateOfBirth time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	Age         int32     `gorm:"column:age" json:"age"`
	Gender      string    `gorm:"column:gender" json:"gender"`
	Role        string    `gorm:"column:role;not null" json:"role"`
	IsBlocked   bool      `gorm:"column:is_blocked" json:"is_blocked"`
	CreatedAt   time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
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

type DriverLocation struct {
	DriverID  string    `gorm:"column:driver_id;primaryKey" json:"driver_id"`
	Location  string    `gorm:"column:location" json:"location"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}
