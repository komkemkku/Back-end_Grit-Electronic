package requests

// type OrderRequest struct {
// 	Page   int64  `form:"page"`
// 	Size   int64  `form:"size"`
// 	Search string `form:"search"`

// }
type OrderRequest struct {
	Page   int64  `json:"page"`
	Size   int64  `json:"size"`
	Search string `json:"search"`
	Status string `json:"status"` // ฟิลด์สำหรับค้นหาตาม status
}

type OrderIdRequest struct {
	ID int `uri:"id"`
}

type OrderCreateRequest struct {
	UserID     int `json:"user_id"`
	PaymentID  int `json:"payment_id"`
	ShipmentID int `json:"shipment_id"`
	// CartID     int    `json:"cart_id"`
	// Status     string `json:"status"`
}

type OrderUpdateRequest struct {
	// ID         int    `json:"id"`
	// UserID     int    `json:"user_id"`
	// PaymentID  int    `json:"payment_id"`
	// ShipmentID int    `json:"shipment_id"`
	// CartID     int    `json:"cart_id"`
	Status string `json:"status"`
}
