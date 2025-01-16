package requests

type OrderDetailRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderDetailIdRequest struct {
	ID int `uri:"id"`
}

type OrderDetailCreateRequest struct {
	OrderID          int     `json:"order_id"`
	ProductID        int     `json:"product_id"`
	ProductName      string  `json:"product_name"`
	ProductImageMain string  `json:"product_image_main"`
	PricePerPerduct  float64 `json:"price_per_producer"`
	AmountPerProduct int     `json:"amount_per_producer"`
}

type OrderDetailUpdateRequest struct {
	ID               string  `json:"id"`
	OrderID          int     `json:"order_id"`
	ProductID        int     `json:"product_id"`
	ProductName      string  `json:"product_name"`
	ProductImageMain string  `json:"product_image_main"`
	PricePerPerduct  float64 `json:"price_per_producer"`
	AmountPerProduct int     `json:"amount_per_producer"`
}
