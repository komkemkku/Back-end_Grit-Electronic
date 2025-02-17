package requests

type WishlistsRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
	UserID int    `form:"user_id"`
}

type WishlistsIdRequest struct {
	ID int `uri:"id"`
}

type WishlistsAddRequest struct {
	UserID           int     `json:"user_id"`
	ProductID        int     `json:"product_id"`
	PricePerProduct  float64 `json:"price_per_product"`
	AmountPerProduct int     `json:"amount_per_product"`
}

type WishlistsUpdateRequest struct {
	ID               int     `json:"id"`
	UserID           int     `json:"user_id"`
	ProductID        int     `json:"product_id"`
}
