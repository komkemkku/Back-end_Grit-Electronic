package requests

type CartRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CartIdRequest struct {
	ID int64 `uri:"id"`
}

type CartAddItemRequest struct {
	UserID    int64 `json:"user_id"`
	ProductID int64 `json:"product_id"`
	Quantity  int `json:"quantity"`
}
