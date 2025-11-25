package models

import (
	"time"
)

type Authentication struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Role      string    `gorm:"column:role" json:"role"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;foreignKey:ID;references:ID"`
	Driver    Driver    `gorm:"constraint:OnDelete:CASCADE;foreignKey:ID;references:ID"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
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

type Order struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID       int32     `gorm:"column:user_id" json:"user_id"`
	DriverID     int32     `gorm:"column:driver_id" json:"driver_id"`
	PickupPoint  string    `gorm:"column:pickup_point" json:"pickup_point"`
	DropoffPoint string    `gorm:"column:dropoff_point" json:"dropoff_point"`
	Status       string    `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
}
