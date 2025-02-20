package requests

type OrderRequest struct {
	Page      int64  `form:"page"`
	Size      int64  `form:"size"`
	Search    string `form:"search"`
	Status    string `form:"status"`
	StartDate int64  `form:"start"`
	EndDate   int64  `form:"end"`
}

type OrderUserRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
	UserID int    `form:"user_id"`
}

type OrderIdRequest struct {
	ID int `uri:"id"`
}

type OrderCreateRequest struct {
	UserID       int         `json:"user_id"`
	ShipmentID   int         `json:"shipment_id"`
	SystemBankID int         `json:"system_bank_id"`
	PaymentDate  string      `json:"payment_date"`
	CartItems    []CartItems `json:"cart_items"`
}

type CartItems struct {
	CartID             int64 `json:"cart_id"`
	TotalProductAmount int64 `json:"total_product_amount"`
}

type OrderUpdateRequest struct {
	ID             int    `json:"id"`
	ShipmentID     int    `json:"shipment_id"`
	Status         string `json:"status"`
	TrackingNumber string `json:"tracking_number"`
}
