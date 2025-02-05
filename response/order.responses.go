package response

type OrderResponses struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	PaymentID   int     `json:"payment_id"`
	ShipmentID  int     `json:"shipment_id"`
	TotalPrice  float64 `json:"total_price"`
	TotalAmount int     `json:"total_amount"`
	Status      string  `json:"status"`
	Created_at  int64   `json:"created_at"`
	Updated_at  int64   `json:"updated_at"`
}

type OrderRespOrderDetail struct {
	ID              int                     `json:"id"`
	User            UserRespOrderDetail     `bun:"user"`
	Payment         PaymentRespOrderDetail  `bun:"payment"`
	SystemBank      SystemBankRespPayment   `bun:"systembank"`
	ImageSystemBank ImageSystemBankResp     `bun:"imagesystembank"`
	Shipment        ShipmentRespOrderDetail `bun:"shipment"`
	TotalAmount     int                     `json:"total_amount"`
	TotalPrice      float64                 `json:"total_price"`
	Status          string                  `json:"status"`
}

// type OrderResponses struct {
// 	ID          int     `json:"id"`
// 	UserID      int     `json:"user_id"`
// 	Username    string  `json:"username"`
// 	Status      string  `json:"status"`
// 	CreatedAt   string  `json:"created_at"`
// 	UpdatedAt   string  `json:"updated_at"`
// 	TotalAmount int     `json:"total_amount"`
// 	TotalPrice  float64 `json:"total_price"`

// 	// ข้อมูลการชำระเงิน
// 	PaymentID     int     `json:"payment_id" bun:"payment_id"`
// 	SystemBankID  int     `json:"system_bank_id" bun:"system_bank_id"`
// 	PaymentPrice  float64 `json:"payment_price" bun:"payment_price"`
// 	BankName      string  `json:"bank_name" bun:"bank_name"`
// 	AccountName   string  `json:"account_name" bun:"account_name"`
// 	AccountNumber string  `json:"account_number" bun:"account_number"`
// 	PaymentStatus string  `json:"payment_status" bun:"payment_status"`

// 	// ข้อมูลการจัดส่ง
// 	ShipmentID     int    `json:"shipment_id" bun:"shipment_id"`
// 	Firstname      string `json:"firstname" bun:"firstname"`
// 	Lastname       string `json:"lastname" bun:"lastname"`
// 	Address        string `json:"address" bun:"address"`
// 	ZipCode        string `json:"zip_code" bun:"zip_code"`
// 	SubDistrict    string `json:"sub_district" bun:"sub_district"`
// 	District       string `json:"district" bun:"district"`
// 	Province       string `json:"province" bun:"province"`
// 	ShipmentStatus string `json:"shipment_status" bun:"shipment_status"`
// }
