package requests

type CartRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CartIdRequest struct {
	ID int `uri:"id"`
}

type CartAddItemRequest struct {
	UserID int `json:"user_id"`
}

type CartUpdateItemRequest struct {
	ID              int     `json:"id"`
	UserID          int     `json:"user_id"`
	TotalCartAmount int     `json:"total_cart_amount"`
	TotalCartPrice  float32 `json:"total_cart_price"`
	Status          string  `json:"status"`
}
