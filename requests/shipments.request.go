package requests

type ShipmentRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ShipmentIdRequest struct {
	ID int64 `uri:"id"`
}

type ShipmentCreateRequest struct {
	Address      string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District     string `json:"district"`
	Province     string `json:"province"`
	UserID      string `json:"user_id"`
	OrderID     string `json:"order_id"`
	Status       string `json:"status"`
}

type ShipmentUpdateRequest struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District     string `json:"district"`
	Province     string `json:"province"`
	UserID      string `json:"user_id"`
	OrderID     string `json:"order_id"`
	Status       string `json:"status"`
}
