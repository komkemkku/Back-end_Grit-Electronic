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
	ProductID  int `json:"product_id"`
	OrderID    int `json:"order_id"`
	PaymentID  int `json:"payment_id"`
	ShipmentID int `json:"shipment_id"`
}

type OrderDetailUpdateRequest struct {
	Id         string `json:"id"`
	PaymentID  int    `json:"payment_id"`
	ShipmentID int    `json:"shipment_id"`
	ProductID  int    `json:"product_id"`
	OrderID    int    `json:"order_id"`
}
