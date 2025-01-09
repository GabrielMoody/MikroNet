package model

import "time"

type UserDetails struct {
	ID          string    `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	FirstName   string    `gorm:"column:first_name;not null;type:varchar(255)" json:"first_name"`
	LastName    string    `gorm:"column:last_name;type:varchar(255)" json:"last_name"`
	Email       string    `gorm:"column:email;type:varchar(255)" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(255)" json:"phone_number"`
	DateOfBirth time.Time `gorm:"column:date_of_birth;type:date" json:"date_of_birth"`
	Age         int32     `gorm:"column:age" json:"age"`
	Gender      string    `gorm:"column:gender;type:varchar(255)" json:"gender"`
}

type Review struct {
	ID       int    `gorm:"column:id;primaryKey" json:"id"`
	UserID   string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	DriverID string `gorm:"column:driver_id;type:varchar(255)" json:"driver_id"`
	Comment  string `gorm:"column:comment;type:varchar(255)" json:"comment"`
	Star     int    `gorm:"column:star" json:"star"`
}
