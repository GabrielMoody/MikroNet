package dto

type (
	StatusReq struct {
		Status string `json:"status" validate:"required"`
	}

	GetDriverDetailsRes struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		LicenseNumber  string `json:"license_number" form:"license_number"`
		SIM            string `json:"sim" form:"sim"`
		ProfilePicture string `json:"profile_picture"`
	}

	EditDriverReq struct {
		Name           string `json:"name" form:"name" validate:"required"`
		LicenseNumber  string `json:"license_number" form:"license_number" validate:"required"`
		SIM            string `json:"sim" form:"sim" validate:"required"`
		ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	}

	LocationReq struct {
		DriverID string  `json:"driver-id"`
		RouteId  int     `json:"route_id"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
	}
)
