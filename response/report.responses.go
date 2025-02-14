package response

type DashboardResponse struct {
	TotalSales     float64 `json:"totalsale"`
	TotalOrders    int     `json:"totalorder"`
	TotalUsers     int     `json:"totaluser"`
	TotalCancelled int     `json:"totalcancelled"`
}

type ProductSales struct {
    ProductName string  `json:"product_name"`
    TotalSales  float64 `json:"total_sales"`
}
type DashboardCategoryResponses struct {
    Category string         `json:"category"`
    Year     int            `json:"year"`
    Month    string         `json:"month"`
    Products []ProductSales `json:"products"`
}

type ReportReponses struct {
	OrderID     int     `bun:"order_id"`
	UserName    string  `bun:"username"`
	ProductName string  `bun:"product_name"`
	Amount      int     `bun:"amount"`
	Price       float64 `bun:"price"`
	TotalPrice  float64 `bun:"total_price"`
	Created_at  int64   `bun:"created_at"`
}
