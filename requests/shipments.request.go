package requests

type ShipmentRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ShipmentIdRequest struct {
	ID int `uri:"id"`
}

type ShipmentCreateRequest struct {
	UserID      int    `json:"user_id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Status      string `json:"status"`
}

type ShipmentUpdateRequest struct {
	ID          string `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Status      string `json:"status"`
}
