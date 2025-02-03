package requests

type CartItemRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CartItemIdRequest struct {
	ID int `uri:"id"`
}

type CartItemCreateRequest struct {
	UserID             int `json:"user_id"`
	ProductID          int `json:"product_id"`
	TotalProductAmount int `json:"total_product_amount"`
	Status string `json:"status"`
}

type CartItemUpdateRequest struct {
	ID                 int    `json:"id"`
	UserID             int    `json:"user_id"`
	ProductID          int    `json:"product_id"`
	// CartID             int    `json:"cart_id"`
	TotalProductAmount int    `json:"total_product_amount"`
	Status             string `json:"status"`
}

type CartItemDeleteRequest struct {
	// CartID     int `json:"cart_id"`
	UserID     int `json:"user_id"`
	CartItemID int `json:"cart_item_id"`
}
