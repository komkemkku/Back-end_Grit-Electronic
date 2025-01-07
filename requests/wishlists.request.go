package requests

type WishlistsRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type WishlistsIdRequest struct {
	ProductID int64 `uri:"product_id"`
}

type WishlistsAddRequest struct {
	ProductID int64 `json:"product_id"`
}

type WishlistsUpdateRequest struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
}
