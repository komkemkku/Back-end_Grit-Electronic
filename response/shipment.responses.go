package response

type ShipmentResponses struct {
	ID          int64  `json:"id"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Status      string `json:"status"`
	Created_at  int64  `json:"created_at"`
	Updated_at  int64  `json:"updated_at"`
}
