package model

import "time"

type DriverLocation struct {
	DriverID  string    `gorm:"column:driver_id;primaryKey" json:"driver_id"`
	RouteID   int       `gorm:"column:route_id" json:"route_id"`
	Location  string    `gorm:"column:location" json:"location"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type DriverLocationLogs struct {
	DriverID  string    `gorm:"column:driver_id" json:"driver_id"`
	RouteID   int       `gorm:"column:route_id" json:"route_id"`
	Location  string    `gorm:"column:location" json:"location"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}
