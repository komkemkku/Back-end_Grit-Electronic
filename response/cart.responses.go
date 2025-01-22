package response

type CartResponses struct {
	ID              int           `json:"id"`
	UserID          int           `json:"user_id"`
	TotalCartAmount int           `json:"total_cart_amount"`
	TotalCartPrice  float32       `json:"total_cart_price"`
	CartItems       []CartItemRes `bun:"cart_items"`
	Status          string        `json:"status"`
	Created_at      int64         `json:"created_at"`
	Updated_at      int64         `json:"updated_at"`
}

type CartItemRes struct {
	ID                 int     `json:"id"`
	ProductID          int     `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductImageMain   string  `json:"product_image_main"`
	TotalProductPrice  float64 `json:"total_product_price"`
	TotalProductAmount int     `json:"total_product_amount"`
}
