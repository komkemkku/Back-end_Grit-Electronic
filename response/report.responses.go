package response

type DashboardResponse struct {
	TotalSales     float64 `json:"totalsale"`
	TotalOrders    int     `json:"totalorder"`
	TotalUsers     int     `json:"totaluser"`
	TotalCancelled int     `json:"totalcancelled"`
}

type DashboardCategoryResponses struct {
	Category   string  `json:"category"`
	TotalSales float64 `json:"totalsales"`
}
