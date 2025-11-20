package model

import "time"

type Drivers struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	LicenseNumber  string `json:"license_number"`
	SIM            string `json:"sim"`
	Verified       bool   `json:"verified"`
	ProfilePicture string `json:"profile_picture"`
}

type Histories struct {
	ID            int       `json:"id"`
	PassengerName string    `json:"passenger_name"`
	DriverName    string    `json:"driver_name"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
