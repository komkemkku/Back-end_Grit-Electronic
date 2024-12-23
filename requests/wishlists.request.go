package requests

type WishlistRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type WishlistIdRequest struct {
	ProductId int64 `uri:"product_id"`
}

type WishlistAddRequest struct {
	Product_id int64 `json:"product_id"`
	User_id    int64 `json:"user_id"`
}

type WishlistUpdateRequest struct {
	Id         int64 `json:"id"`
	Product_id int64 `json:"product_id"`
	User_id    int64 `json:"user_id"`
}
