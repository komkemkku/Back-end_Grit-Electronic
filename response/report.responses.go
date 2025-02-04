package response

type DashboardResponse struct {
	TotalSales   float64 `json:"totalsale"`
	TotalOrders  int     `json:"totalorder"`
	TotalUsers   int     `json:"totaluser"`
	TotalRefunds int     `json:"totalrefunds"`
}
