package requests

type ShipmentRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ShipmentIdRequest struct {
	Id int64 `uri:"id"`
}

type ShipmentCreateRequest struct {
	Address      string `json:"address"`
	Zip_code     string `json:"zip_code"`
	Sub_district string `json:"sub_district"`
	District     string `json:"district"`
	Province     string `json:"province"`
	User_id      string `json:"user_id"`
	Order_id     string `json:"order_id"`
	Status       string `json:"status"`
}

type ShipmentUpdateRequest struct {
	Id           string `json:"id"`
	Address      string `json:"address"`
	Zip_code     string `json:"zip_code"`
	Sub_district string `json:"sub_district"`
	District     string `json:"district"`
	Province     string `json:"province"`
	User_id      string `json:"user_id"`
	Order_id     string `json:"order_id"`
	Status       string `json:"status"`
}
