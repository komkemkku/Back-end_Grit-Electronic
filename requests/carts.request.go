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
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartUpdateItemRequest struct {
	ID        int `json:"id"`
	ProductID int `json:"product"`
	Quantity  int `json:"quantity"`
}
