package model

import (
	"time"
)

type DriverDetails struct {
	ID             string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	OwnerID        string `gorm:"column:owner_id;type:varchar(255)" json:"owner_id"`
	Email          string `gorm:"column:email;type:varchar(255)" json:"email"`
	Name           string `gorm:"column:name;type:varchar(255)" json:"name"`
	PhoneNumber    string `gorm:"column:phone_number;type:varchar(255)" json:"phone_number"`
	RouteID        *uint  `gorm:"foreignKey:RouteID;references:ID;"`
	Route          Route
	LicenseNumber  string `gorm:"column:license_number;type:varchar(255)" json:"license_number" form:"license_number"`
	SIM            string `gorm:"column:sim;type:varchar(255)" json:"sim" form:"sim"`
	Status         string `gorm:"column:status;type:varchar(255)" json:"status" form:"status"`
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
