package requests

type OrderDetailRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderDetailIdRequest struct {
	Id int64 `uri:"id"`
}

type OrderDetailCreateRequest struct {
	Quantity    int64   `json:"quantity"`
	Product_id  int64   `json:"product_id"`
	Order_id    int64   `json:"order_id"`
	Unit_price  float64 `json:"unit_price"`
	Total_price int64   `json:"total_price"`
}

type OrderDetailUpdateRequest struct {
	Id          string  `json:"id"`
	Quantity    int64   `json:"quantity"`
	Product_id  int64   `json:"product_id"`
	Order_id    int64   `json:"order_id"`
	Unit_price  float64 `json:"unit_price"`
	Total_price int64   `json:"total_price"`
}
