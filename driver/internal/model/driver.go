package model

import (
	"time"
)

type DriverDetails struct {
	ID             string    `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	OwnerID        string    `gorm:"column:owner_id;type:varchar(255)" json:"owner_id"`
	Email          string    `gorm:"column:email;type:varchar(255)" json:"email"`
	FirstName      string    `gorm:"column:first_name;type:varchar(255)" json:"first_name"`
	LastName       string    `gorm:"column:last_name;type:varchar(255)" json:"last_name"`
	DateOfBirth    time.Time `gorm:"column:date_of_birth;type:DATE" json:"date_of_birth"`
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(255)" json:"phone_number"`
	Age            int32     `gorm:"column:age" json:"age"`
	Gender         string    `gorm:"column:gender;type:varchar(255)" json:"gender"`
	RouteID        *uint     `gorm:"foreignKey:RouteID;references:ID;"`
	Route          Route
	LicenseNumber  string `gorm:"column:license_number;type:varchar(255)" json:"license_number" form:"license_number"`
	Status         string `gorm:"column:status;type:varchar(255)" json:"status" form:"status"`
	AvailableSeats int32  `gorm:"column:available_seats" json:"available_seats" form:"available_seats"`
	Verified       bool   `gorm:"column:verified;default:false" json:"verified" form:"verified"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profile_picture" form:"photo"`
}

type Route struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	RouteName string `gorm:"column:route_name;type:varchar(255)" json:"route_name"`
}

type Trip struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	DriverID    string    `gorm:"column:driver_id" json:"driver_id"`
	Location    string    `gorm:"column:location" json:"location"`
	Destination string    `gorm:"column:destination" json:"destination"`
	TripDate    time.Time `gorm:"column:trip_date" json:"trip_date"`
}
