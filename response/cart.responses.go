package response

type CartResponses struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	ProductID        int     `json:"product_id"`
	PricePerProduct  float32 `json:"price_per_product"`
	AmountPerProduct int     `json:"amount_per_product"`
	TotalCartAmount  int     `json:"total_cart_amount"`
	TotalCartPrice   float32 `json:"total_cart_price"`
	Status           string  `json:"status"`
	Created_at       int64   `json:"created_at"`
	Updated_at       int64   `json:"updated_at"`
}
