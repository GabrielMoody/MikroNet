package model

import (
	"time"
)

type User struct {
	ID              string `gorm:"primaryKey;type:varchar(255)"`
	Email           string `gorm:"unique;type:varchar(255)"`
	Password        string
	Role            string `gorm:"type:enum('admin','user','driver')"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DriverDetail    DriverDetails    `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	PassengerDetail PassengerDetails `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	AdminDetail     Admin            `gorm:"foreignKey:ID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type DriverDetails struct {
	ID             string `gorm:"type:varchar(255);primaryKey"`
	Name           string `gorm:"type:varchar(255)"`
	PhoneNumber    string `gorm:"type:varchar(255)"`
	RouteID        *uint
	Route          Route  `gorm:"foreignKey:RouteID;references:ID"`
	LicenseNumber  string `gorm:"type:varchar(255)"`
	SIM            string `gorm:"type:varchar(255)"`
	Status         string `gorm:"type:varchar(255)"`
	Verified       bool   `gorm:"default:false"`
	AvailableSeats int
	QrisData       string
	ProfilePicture string `gorm:"type:varchar(255)"`
	KTP            string `gorm:"type:varchar(255)"`
}

type PassengerDetails struct {
	ID   string `gorm:"primaryKey;type:varchar(255)"`
	Name string `gorm:"type:varchar(255)"`
}

type Admin struct {
	ID   string `gorm:"primaryKey;type:varchar(255)"`
	Name string `gorm:"type:varchar(255)"`
}

type ResetPassword struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"unique;type:varchar(255)"`
	User   User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Code   string `gorm:"type:varchar(255)"`
}

type BlockedAccount struct {
	ID     int    `gorm:"primaryKey"`
	UserID string `gorm:"type:varchar(255);unique"`
	User   User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Route struct {
	ID        uint   `gorm:"primaryKey"`
	RouteName string `gorm:"type:varchar(255)"`
}

type Review struct {
	ID          int              `gorm:"primaryKey"`
	PassengerID string           `gorm:"type:varchar(255)"`
	Passenger   PassengerDetails `gorm:"foreignKey:PassengerID;references:ID;constraint:OnDelete:CASCADE"`
	DriverID    string           `gorm:"type:varchar(255)"`
	Driver      DriverDetails    `gorm:"foreignKey:DriverID;references:ID;constraint:OnDelete:CASCADE"`
	Comment     string           `gorm:"type:varchar(255)"`
	Star        int
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Transaction struct {
	ID          int       `gorm:"primaryKey"`
	PassengerID string    `gorm:"type:varchar(255)"`
	Passenger   User      `gorm:"foreignKey:PassengerID;references:ID;constraint:OnDelete:CASCADE"`
	DriverID    string    `gorm:"type:varchar(255)"`
	Driver      User      `gorm:"foreignKey:DriverID;references:ID;constraint:OnDelete:CASCADE"`
	Amount      int       `gorm:"type:int"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
