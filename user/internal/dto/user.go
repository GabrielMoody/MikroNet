package dto

type (
	EditUserDetails struct {
		FirstName   string `json:"first_name" validate:"required"`
		LastName    string `json:"last_name" validate:"required"`
		DateOfBirth string `json:"date_of_birth"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
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
)
