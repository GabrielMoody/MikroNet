package dto

type (
	EditUserDetails struct {
		Name string `json:"name" validate:"required"`
	}

	MessageLoc struct {
		UserID int64   `json:"user_id"`
		Lat    float64 `json:"lat"`
		Lng    float64 `json:"lng"`
	}

	Point struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	OrderReq struct {
		UserID      int64 `json:"user_id"`
		PickupPoint Point `json:"pickup_point"`
		DestPoint   Point `json:"dest_point"`
	}

	OrderNotificationData struct {
		RecipientID string `json:"recipent_id"`
		Title       string `json:"title"`
		Passenger   string `json:"passenger"`
		PickupPoint Point  `json:"pickup_point"`
		DestPoint   Point  `json:"dest_point"`
	}
)
