package dto

import "time"

type (
	StatusReq struct {
		Status string `json:"status"`
	}

	SeatReq struct {
		Seat int32 `json:"seat"`
	}

	EditDriverReq struct {
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		DateOfBirth time.Time `json:"date_of_birth"`
		Age         int       `json:"age"`
		Gender      string    `json:"gender"`
	}

	LocationReq struct {
		DriverID string  `json:"driver-id"`
		RouteId  int     `json:"route_id"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
	}
)
