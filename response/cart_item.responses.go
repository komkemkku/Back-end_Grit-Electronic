package response

type CartItemResponses struct {
	ID                 int                 `json:"id"`
	CartID             int                 `json:"cart_id"`
	Product            ProductRespCartItem `bun:"product"`
	ProductName        string              `json:"product_name"`
	ProductImageMain   string              `json:"product_image_main"`
	TotalProductPrice  float32             `json:"total_product_price"`
	TotalProductAmount int                 `json:"total_product_amount"`
	Status             string              `json:"status"`
	Created_at         int64               `json:"created_at"`
	Updated_at         int64               `json:"updated_at"`
}
