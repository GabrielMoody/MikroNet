package dto

type (
	EditUserDetails struct {
		Name string `json:"name" validate:"required"`
	}

	MessageLoc struct {
		UserID string  `json:"user_id"`
		Lat    float64 `json:"lat"`
		Lng    float64 `json:"lng"`
	}

	ReviewReq struct {
		Comment string `json:"comment"`
		Star    int    `json:"star" validate:"required"`
	}

	Transaction struct {
		DriverId string `json:"driver_id"`
		Amount   int    `json:"amount"`
	}
)
