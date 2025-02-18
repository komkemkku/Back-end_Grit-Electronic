package response

type UserResponses struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
	Shipment  *Shipment `json:"shipment,omitempty"`
}

type Shipment struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Address     string `json:"address"`
	ZipCode     string `json:"zip_code"`
	SubDistrict string `json:"sub_district"`
	District    string `json:"district"`
	Province    string `json:"province"`
}

type UserRespReview struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

type UserRespWistlist struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserRespCart struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserRespShipment struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserRespOrderDetail struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
}
