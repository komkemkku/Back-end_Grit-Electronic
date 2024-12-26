package requests

type OrderDetailRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderDetailIdRequest struct {
	ID int64 `uri:"id"`
}

type OrderDetailCreateRequest struct {
	Quantity    int64   `json:"quantity"`
	ProductID   int64   `json:"product_id"`
	OrderID     int64   `json:"order_id"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice int64   `json:"total_price"`
}

type OrderDetailUpdateRequest struct {
	Id          string  `json:"id"`
	Quantity    int64   `json:"quantity"`
	ProductID   int64   `json:"product_id"`
	OrderID     int64   `json:"order_id"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice int64   `json:"total_price"`
}
