package requests

type OrderRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type OrderIdRequest struct {
	Id int64 `uri:"id"`
}

type OrderCreateRequest struct {
	User_id      int64  `json:"user_id"`
	Shipment_id  int64  `json:"shipment_id"`
	Payment_id   int64  `json:"payment_id"`
	Total_price  int64  `json:"total_price"`
	Total_amount int64  `json:"total_amount"`
	Status       string `json:"status"`
}

type OrderUpdateRequest struct {
	Id           int64  `json:"id"`
	User_id      int64  `json:"user_id"`
	Shipment_id  int64  `json:"shipment_id"`
	Payment_id   int64  `json:"payment_id"`
	Total_price  int64  `json:"total_price"`
	Total_amount int64  `json:"total_amount"`
	Status       string `json:"status"`
}
