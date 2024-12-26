package response

type OrderResponses struct {
	ID           int                    `json:"id"`
	Total_Price  float64                `json:"total_price"`
	Total_Amount int64                  `json:"total_amount"`
	Status       string                 `json:"status"`
	User_id      int64                  `json:"user_id"`
	Shipment_id  int64                  `json:"shipment_id"`
	Payment_id   int64                  `json:"payment_id"`
	Created_at   int64                  `json:"created_at"`
	Updated_at   int64                  `json:"updated_at"`
	OrderDetails []OrderDetailResponses `json:"order_details"`
}

type OrderDetailResponses struct {
	ID         int64                `json:"id"`
	Product_id int64                `json:"product_id"`
	Quantity   int64                `json:"quantity"`
	Price      float64              `json:"price"`
	Product    ProductListResponses `json:"product"`
}

type ProductListResponses struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}
