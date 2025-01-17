package response

// type WishlistResponses struct {
// 	ID         int             `json:"id"`
// 	Product    ProductRespCart `json:"product"`
// 	Created_at int64           `json:"created_at"`
// 	Updated_at int64           `json:"updated_at"`
// }

type WishlistResponses struct {
	ID               int              `json:"id"`
	UserID           UserRespWistlist `json:"user"`
	ProductID        ProductRespCart  `json:"product"`
	PricePerProduct  float64          `json:"price_per_product"`
	AmountPerProduct int              `json:"amount_per_product"`
	CreatedAt        int64            `json:"created_at"`
	UpdatedAt        int64            `json:"updated_at"`
}
