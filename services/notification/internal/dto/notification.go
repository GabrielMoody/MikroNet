package dto

type Geo struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type OrderNotificationData struct {
	RecipientID string `json:"recipent_id"`
	Title       string `json:"title"`
	Passenger   string `json:"passenger"`
	PickupPoint Geo    `json:"pickup_point"`
	DestPoint   Geo    `json:"dest_point"`
}
