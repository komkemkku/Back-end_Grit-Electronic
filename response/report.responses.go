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
	Quantity    int     `json:"quantity"`
}

type DashboardCategoryResponses struct {
	Category           string         `json:"category"`
	TotalCategorySales float64        `json:"total_category_sales"`
	Products           []ProductSales `json:"products"`
	Month              string         `json:"month"`
	Year               string         `json:"year"`
}

type ReportReponses struct {
	OrderID     int            `json:"order_id" bun:"order_id"`
	UserName    string         `json:"username" bun:"username"`
	Firstname   string         `json:"firstname" bun:"firstname"`
	Lastname    string         `json:"lastname" bun:"lastname"`
	TotalAmount int            `json:"total_amount" bun:"total_amount"`
	TotalPrice  float64        `json:"total_price" bun:"total_price"`
	Products    []ProductInfoo `json:"products" bun:"products"`
	Created_at  int64          `json:"created_at" bun:"created_at"`
}

type ProductInfoo struct {
	ProductName        string  `json:"product_name"`
	Price              float64 `json:"price"`
	TotalProductAmount int     `json:"total_product_amount"`
}
