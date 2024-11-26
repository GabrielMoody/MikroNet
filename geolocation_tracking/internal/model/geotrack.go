package model

import "time"

//type Driver struct {
//	ID                 string `gorm:"column:id;primaryKey" json:"id"`
//	OwnerID            string `gorm:"column:owner_id" json:"owner_id"`
//	RegistrationNumber string `gorm:"column:registration_number" json:"registration_number"`
//	Status             string `gorm:"column:status" json:"status"`
//	Location           string `gorm:"column:location" json:"location"`
//	AvailableSeats     int32  `gorm:"column:available_seats" json:"available_seats"`
//}

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
