package models

import "time"

type User struct {
	ID           string `gorm:"primaryKey"`
	NamaLengkap  string
	Email        string `gorm:"unique"`
	Password     string
	NomorTelepon string `gorm:"unique"`
	JenisKelamin string
	TanggalLahir *time.Time `gorm:"type:date"`
	ImageUrl     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
