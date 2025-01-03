package dto

type (
	EditUserDetails struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		DateOfBirth string `json:"date_of_birth"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
	}

	OrderReq struct {
		From      string `json:"from" validate:"required"`
		To        string `json:"to" validate:"required"`
		Passenger int    `json:"passenger" validate:"required"`
	}

	CarterReq struct {
		PickUpPoint string `json:"pick_up_point" validate:"required"`
		PickUpTime  string `json:"pick_up_time" validate:"required"`
		Duration    string `json:"duration" validate:"required"`
		Comments    string `json:"comments,omitempty"`
	}

	Orders struct {
		Id                 string
		FirstName          string
		LastName           string
		RegistrationNumber string
		Distance           string
	}

	CurrLocation struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	}

	ReviewReq struct {
		Comment string `json:"comment"`
		Star    int    `json:"star" validate:"required"`
	}
)
