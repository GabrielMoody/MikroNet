package dto

type (
	StatusReq struct {
		Status string `json:"status"`
	}

	SeatReq struct {
		Seat int32 `json:"seat"`
	}
)
