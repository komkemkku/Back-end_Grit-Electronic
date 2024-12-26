package requests

type OrderRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderIdRequest struct {
	ID int64 `uri:"id"`
}

type OrderCreateRequest struct {
	UserID       int64  `json:"user_id"`
	ShipmentID  int64  `json:"shipment_id"`
	PaymentID   int64  `json:"payment_id"`
	TotalPrice  int64  `json:"total_price"`
	TotalAmount int64  `json:"total_amount"`
	Status       string `json:"status"`
}

type OrderUpdateRequest struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	ShipmentID  int64  `json:"shipment_id"`
	PaymentID   int64  `json:"payment_id"`
	TotalPrice  int64  `json:"total_price"`
	TotalAmount int64  `json:"total_amount"`
	Status       string `json:"status"`
}
