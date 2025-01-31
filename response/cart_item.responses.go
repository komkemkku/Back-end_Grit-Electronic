package response

type CartItemResponses struct {
	ID                 int                 `json:"id"`
	CartID             int                 `json:"cart_id"`
	Product            ProductRespCartItem `bun:"product"`
	TotalProductAmount int                 `json:"total_product_amount"`
	Status             string              `json:"status"`
	Created_at         int64               `json:"created_at"`
	Updated_at         int64               `json:"updated_at"`
}

type CartItemRes struct {
	ID                 int                 `json:"id"`
	Product            ProductRespCartItem `json:"product"`
	TotalProductAmount int                 `json:"total_product_amount"`
}
