package response


type OrderResponses struct {
    ID         int    `json:"id"`
    UserID     int    `json:"user_id"`
    PaymentID  int    `json:"payment_id"`
    ShipmentID int    `json:"shipment_id"`
    CartID     int    `json:"cart_id"`
    Status     string `json:"status"`
    Created_at int64  `json:"created_at"`
    Updated_at int64  `json:"updated_at"`
}


// type OrderResponses struct {
// 	ID          int     `json:"id"`
// 	TotalPrice  float64 `json:"total_price"`
// 	TotalAmount int     `json:"total_amount"`
// 	Status      int     `json:"status"`
// 	Created_at  int64   `json:"created_at"`
// 	Updated_at  int64   `json:"updated_at"`
// }

type OrderRespOrderDetail struct {
	ID          int     `json:"id"`
	Status      int     `json:"status"`
}

// TotalPrice  float64 `json:"total_price"`
// TotalAmount int     `json:"total_amount"`
