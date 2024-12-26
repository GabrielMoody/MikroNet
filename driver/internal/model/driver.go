package model

import (
	"time"
)

type DriverDetails struct {
	ID                 string    `gorm:"column:id;primaryKey" json:"id"`
	OwnerID            string    `gorm:"column:owner_id" json:"owner_id"`
	Email              string    `gorm:"column:email" json:"email"`
	FirstName          string    `gorm:"column:first_name" json:"first_name"`
	LastName           string    `gorm:"column:last_name" json:"last_name"`
	DateOfBirth        time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	Age                int32     `gorm:"column:age" json:"age"`
	Gender             string    `gorm:"column:gender" json:"gender"`
	RouteID            int       `gorm:"column:route_id" json:"route_id"`
	Route              Route     `gorm:"foreignKey:RouteID;references:ID"`
	RegistrationNumber string    `gorm:"column:registration_number" json:"registration_number"`
	Status             string    `gorm:"column:status" json:"status"`
	AvailableSeats     int32     `gorm:"column:available_seats" json:"available_seats"`
}

type Route struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	RouteName string `gorm:"column:route_name" json:"route_name"`
}

type Trip struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	DriverID    string    `gorm:"column:driver_id" json:"driver_id"`
	Location    string    `gorm:"column:location" json:"location"`
	Destination string    `gorm:"column:destination" json:"destination"`
	TripDate    time.Time `gorm:"column:trip_date" json:"trip_date"`
}

type Review struct {
	ID       string `gorm:"column:id;primaryKey" json:"id"`
	UserID   string `gorm:"column:user_id" json:"user_id"`
	DriverID string `gorm:"column:driver_id" json:"driver_id"`
	Review   string `gorm:"column:review" json:"review"`
	Star     int    `gorm:"column:star" json:"star"`
}
