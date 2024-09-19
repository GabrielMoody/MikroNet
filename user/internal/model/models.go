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
}

type Order struct {
	ID              string `gorm:"column:id;primaryKey" json:"id"`
	UserID          string `gorm:"column:user_id" json:"user_id"`
	DriverID        string `gorm:"column:driver_id" json:"driver_id"`
	PickUpLocation  string `gorm:"column:pickup_location" json:"pickup_location"`
	DropOffLocation string `gorm:"column:drop_off_location" json:"drop_off_location"`
	Status          string `gorm:"column:status" json:"status"`
}

type Route struct {
	ID               string `gorm:"column:id;primaryKey" json:"id"`
	RouteName        string `gorm:"column:route_name" json:"route_name"`
	InitialRoute     string `gorm:"column:initial_route" json:"initial_route"`
	DestinationRoute string `gorm:"column:destination_route" json:"destination_route"`
}

type Review struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	UserID   string `gorm:"column:user_id" json:"user_id"`
	DriverID string `gorm:"column:driver_id" json:"driver_id"`
	Review   string `gorm:"column:review" json:"review"`
	Star     int    `gorm:"column:star" json:"star"`
}
