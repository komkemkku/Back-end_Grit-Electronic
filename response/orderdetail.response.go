package response

type OrderDetailResponses struct {
	ID          int                      `json:"id"`
	TotalPrice  float64                  `json:"total_price"`
	TotalAmount int                      `json:"total_amount"`
	Status      int                      `json:"status"`
	Order       OrderRespOrderDetail     `json:"order"`
	Product     []ProductRespOrderDetail `json:"product"`
	Payment     PaymentRespOrderDetail   `json:"payment"`
	Shipment    ShipmentRespOrderDetail  `json:"shipment"`
	CreatedAt   int64                    `json:"created_at"`
	UpdatedAt   int64                    `json:"updated_at"`
}
