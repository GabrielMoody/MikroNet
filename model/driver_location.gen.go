// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDriverLocation = "driver_location"

// DriverLocation mapped from table <driver_location>
type DriverLocation struct {
	DriverID  string    `gorm:"column:driver_id;primaryKey" json:"driver_id"`
	Location  string    `gorm:"column:location" json:"location"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName DriverLocation's table name
func (*DriverLocation) TableName() string {
	return TableNameDriverLocation
}
