package dto

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type OrderNotificationData struct {
	RecipientID string `json:"recipent_id"`
	Title       string `json:"title"`
	Passenger   string `json:"passenger"`
	OrderID     int    `json:"order_id"`
	PickupPoint Point  `json:"pickup_point"`
	DestPoint   Point  `json:"dest_point"`
}
