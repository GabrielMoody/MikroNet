package model

import (
	"time"
)

type Order struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       *int64    `gorm:"column:user_id" json:"user_id"`
	DriverID     *int64    `gorm:"column:driver_id" json:"driver_id"`
	PickupPoint  GeoPoint  `gorm:"column:pickup_point;type:geometry(Point,4326)" json:"pickup_point"`
	DropoffPoint GeoPoint  `gorm:"column:dropoff_point;type:geometry(Point,4326)" json:"dropoff_point"`
	Status       string    `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
}

type User struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username    string    `gorm:"column:username" json:"username"`
	Fullname    string    `gorm:"column:fullname" json:"fullname"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
}

type Driver struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	VehicleType string    `gorm:"column:vehicle_type" json:"vehicle_type"`
	PlateNumber string    `gorm:"column:plate_number" json:"plate_number"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
}

type DriverStatus struct {
	DriverID       int32      `gorm:"column:driver_id;primaryKey" json:"driver_id"`
	IsOnline       *bool      `gorm:"column:is_online" json:"is_online"`
	IsBusy         *bool      `gorm:"column:is_busy" json:"is_busy"`
	LastActivityAt *time.Time `gorm:"column:last_activity_at" json:"last_activity_at"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
