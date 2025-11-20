package dto

type (
	GetDriverQuery struct {
		Verified *bool `query:"verified"`
	}

	AddRoute struct {
		RouteName string `json:"route_name"`
		Price     int    `json:"price"`
	}

	MonthReport struct {
		Month int `query:"month"`
	}

	RoutesReport struct {
		Route   string `json:"route"`
		Total   int    `json:"total"`
		Revenue int64  `json:"revenue"`
	}

	CommonReport struct {
		TotalPassenger int   `json:"total_passenger"`
		TotalDriver    int   `json:"total_driver"`
		TotalTrip      int   `json:"total_trip"`
		TotalRevenue   int64 `json:"total_revenue"`
	}
	Report struct {
		Common CommonReport   `json:"common"`
		Trips  []RoutesReport `json:"trips"`
	}

	EditAmount struct {
		Amount int `json:"amount"`
	}
)
