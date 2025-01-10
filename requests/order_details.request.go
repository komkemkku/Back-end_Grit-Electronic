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
	Quantity   int     `json:"quantity"`
	ProductID  int     `json:"product_id"`
	OrderID    int     `json:"order_id"`
	PaymentID  int     `json:"payment_id"`
	ShipmentID int     `json:"shipment_id"`
	UnitPrice  float64 `json:"unit_price"`
	//TotalPrice int     `json:"total_price"`
}

type OrderDetailUpdateRequest struct {
	Id         string  `json:"id"`
	Quantity   int     `json:"quantity"`
	ProductID  int     `json:"product_id"`
	OrderID    int     `json:"order_id"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice int     `json:"total_price"`
}
