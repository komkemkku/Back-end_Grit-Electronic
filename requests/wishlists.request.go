package requests

type WishlistsRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type WishlistsIdRequest struct {
	ID int `uri:"id"`
}

type WishlistsAddRequest struct {
	ProductID int `json:"product_id"`
}

type WishlistsUpdateRequest struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
}
