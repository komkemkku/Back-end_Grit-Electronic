package requests

type WishlistRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type WishlistIdRequest struct {
	ProductID int64 `uri:"product_id"`
}

type WishlistAddRequest struct {
	ProductID int64 `json:"product_id"`
	UserID    int64 `json:"user_id"`
}

type WishlistUpdateRequest struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	UserID    int64 `json:"user_id"`
}
