package model

import (
	"database/sql/driver"
	"fmt"
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

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (g GeoPoint) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", g.Lng, g.Lat), nil
}

func (g *GeoPoint) Scan(src interface{}) error {
	var point string

	switch v := src.(type) {
	case []byte:
		point = string(v)
	case string:
		point = v
	default:
		return fmt.Errorf("unsupported type %T", src)
	}

	var lng, lat float64
	_, err := fmt.Sscanf(point, "SRID=4326;POINT(%f %f)", &lng, &lat)
	if err != nil {
		return err
	}

	g.Lat = lat
	g.Lng = lng
	return nil
}
