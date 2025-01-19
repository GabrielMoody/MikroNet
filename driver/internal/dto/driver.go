package dto

import "time"

type (
	StatusReq struct {
		Status string `json:"status" validate:"required"`
	}

	GetDriverDetailsRes struct {
		ID                 string    `json:"id"`
		FirstName          string    `json:"first_name"`
		LastName           string    `json:"last_name"`
		Email              string    `json:"email"`
		DateOfBirth        time.Time `json:"date_of_birth"`
		RegistrationNumber string    `json:"registration_number"`
		Age                int       `json:"age"`
		Gender             string    `json:"gender"`
		ProfilePicture     string    `json:"profile_picture"`
	}

	EditDriverReq struct {
		FirstName      string `json:"first_name" form:"first_name" validate:"required"`
		LastName       string `json:"last_name" form:"last_name" validate:"required"`
		DateOfBirth    string `json:"date_of_birth" form:"date_of_birth"`
		Age            int    `json:"age" form:"age"`
		Gender         string `json:"gender" form:"gender"`
		ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	}

	LocationReq struct {
		DriverID string  `json:"driver-id"`
		RouteId  int     `json:"route_id"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
	}
)
