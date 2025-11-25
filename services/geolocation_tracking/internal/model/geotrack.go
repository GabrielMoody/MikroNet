package model

import "time"

type DriverLocation struct {
	DriverID  int32     `gorm:"column:driver_id;primaryKey" json:"driver_id"`
	Location  string    `gorm:"column:location" json:"location"`
	Heading   float64   `gorm:"column:heading" json:"heading"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
}
