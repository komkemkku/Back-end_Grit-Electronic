package response

import "time"

type OrderResponses struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	Username            string    `json:"username"` // เพิ่มฟิลด์นี้
	UserFirstname       string    `json:"firstname"`
	UserLastname        string    `json:"lastname"`
	UserPhone           string    `json:"phone"`
	TrackingNumber      string    `json:"tracking_number"`
	TotalPrice          float64   `json:"total_price"`
	TotalAmount         int       `json:"total_amount"`
	Status              string    `json:"status"`
	PaymentID           int       `json:"payment_id"`
	ShipmentID          int       `json:"shipment_id"`
	ShipmentFirstname   string    `json:"shipment_firstname"`
	ShipmentLastname    string    `json:"shipment_lastname"`
	ShipmentAddress     string    `json:"shipment_address"`
	ShipmentZipCode     string    `json:"shipment_zip_code"`
	ShipmentSubDistrict string    `json:"shipment_sub_district"`
	ShipmentDistrict    string    `json:"shipment_district"`
	ShipmentProvince    string    `json:"shipment_province"`
	CreatedAt           time.Time `json:"created_at"` // เก็บค่าเป็น int64
	UpdatedAt           time.Time `json:"updated_at"`
}

// OrderRespOrderDetail struct ที่เพิ่มฟิลด์ SystemBankID เพื่อรับข้อมูล system_bank__id
type OrderRespOrderDetail struct {
	ID              int                     `json:"id"`
	User            UserRespOrderDetail     `bun:"user"`
	Products        []string                `json:"product_name"`
	TotalAmount     int                     `json:"total_amount"`
	TotalPrice      float64                 `json:"total_price"`
	Status          string                  `json:"status"`
	TrackingNumber  string                  `json:"tracking_number"`
	Payment         PaymentRespOrderDetail  `bun:"payment"`
	SystemBank      SystemBankRespPayment   `bun:"system_bank"`
	ImageSystemBank ImageSystemBankResp     `bun:"image"`
	Shipment        ShipmentRespOrderDetail `bun:"shipment"`
	Created_at      int64                   `json:"created_at"`
	Updated_at      int64                   `json:"updated_at"`
}

// type OrderProduct struct {
//     ProductName string `json:"product_name"`
//     Amount      int64  `json:"amount"`
//     // หากต้องการฟิลด์อื่นเพิ่มเติม (ราคา, product_id, ฯลฯ) ก็ใส่เพิ่มได้
// }
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
// 	Firstname      string `json:"firstname" bun:"firstname"`
// 	Lastname       string `json:"lastname" bun:"lastname"`
// 	Address        string `json:"address" bun:"address"`
// 	ZipCode        string `json:"zip_code" bun:"zip_code"`
// 	SubDistrict    string `json:"sub_district" bun:"sub_district"`
// 	District       string `json:"district" bun:"district"`
// 	Province       string `json:"province" bun:"province"`
// 	ShipmentStatus string `json:"shipment_status" bun:"shipment_status"`
// }
