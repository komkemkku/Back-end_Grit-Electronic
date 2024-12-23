package requests

type CartRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CartIdRequest struct {
	Id int64 `uri:"id"`
}

type CartAddItemRequest struct {
	User_id    int64 `json:"user_id"`
	Product_id int64 `json:"product_id"`
	Quantity   int64 `json:"quantity"`
}
