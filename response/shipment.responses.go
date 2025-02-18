package response

// type ShipmentResponses struct {
// 	ID          int    `bun:"shipment_id" json:"id"`
// 	Firstname   string `bun:"shipment_firstname" json:"firstname"`
// 	Lastname    string `bun:"shipment_lastname" json:"lastname"`
// 	Address     string `bun:"shipment_address" json:"address"`
// 	ZipCode     string `bun:"shipment_zip_code" json:"zip_code"`
// 	SubDistrict string `bun:"shipment_sub_district" json:"sub_district"`
// 	District    string `bun:"shipment_district" json:"district"`
// 	Province    string `bun:"shipment_province" json:"province"`
// 	Status      string `bun:"shipment_status" json:"status"`
// 	CreatedAt   int64  `bun:"shipment_created_at" json:"created_at"`
// 	UpdatedAt   int64  `bun:"shipment_updated_at" json:"updated_at"`
// }

type ShipmentResponses struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Status      string `json:"status"`
	Created_at  int64  `json:"created_at"`
	Updated_at  int64  `json:"updated_at"`
}

type ShipmentRespOrderDetail struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
}

type ShipmentRespOrder struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
}

type UserAdress struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
}
