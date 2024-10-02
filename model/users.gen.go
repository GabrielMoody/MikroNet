// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	FirstName   string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Email       string    `gorm:"column:email;not null" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number;not null" json:"phone_number"`
	Password    string    `gorm:"column:password;not null" json:"password"`
	DateOfBirth time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	Age         int32     `gorm:"column:age" json:"age"`
	Gender      string    `gorm:"column:gender" json:"gender"`
	Role        string    `gorm:"column:role;not null" json:"role"`
	IsBlocked   bool      `gorm:"column:is_blocked" json:"is_blocked"`
	CreatedAt   time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}