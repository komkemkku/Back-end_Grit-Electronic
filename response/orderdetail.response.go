package response

type OrderDetailResponses struct {
	ID         int   `json:"id"`
	OrderID    int   `json:"order_id"`
	ProductID  int   `json:"product_id"`
	PaymentID  int   `json:"payment_id"`
	ShipmentID int   `json:"shipment_id"`
	Quantity   int   `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	CreatedAt  int64 `json:"created_at"`
	UpdatedAt  int64 `json:"updated_at"`
}
