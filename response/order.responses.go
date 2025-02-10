package response

type OrderResponses struct {
    ID             int     `json:"id"`
    UserID         int     `json:"user_id"`
    Username       string  `json:"username"`
    FirstName      string  `bun:"user_firstname" json:"firstname"`
    LastName       string  `bun:"user_lastname" json:"lastname"`
    PaymentID      int     `json:"payment_id"`
    ShipmentID     int     `bun:"shipment_id" json:"shipment_id"` // ✅ ตรวจสอบให้แน่ใจว่าตรงกัน
    ShipmentFirst  string  `bun:"shipment_firstname" json:"shipment_firstname"`
    ShipmentLast   string  `bun:"shipment_lastname" json:"shipment_lastname"`
    ShipmentAddr   string  `bun:"shipment_address" json:"shipment_address"`
    ShipmentZip    string  `bun:"shipment_zip_code" json:"shipment_zip_code"`
    ShipmentSubDis string  `bun:"shipment_sub_district" json:"shipment_sub_district"`
    ShipmentDis    string  `bun:"shipment_district" json:"shipment_district"`
    ShipmentProv   string  `bun:"shipment_province" json:"shipment_province"`
    TotalPrice     float64 `json:"total_price"`
    TotalAmount    int     `json:"total_amount"`
    Status         string  `json:"status"`
    CreatedAt      int64   `json:"created_at"`
    UpdatedAt      int64   `json:"updated_at"`
  }
  
// OrderRespOrderDetail struct ที่เพิ่มฟิลด์ SystemBankID เพื่อรับข้อมูล system_bank__id
type OrderRespOrderDetail struct {
  ID              int                     `json:"id"`
  User            UserRespOrderDetail     `bun:"user"`
  Payment         PaymentRespOrderDetail  `bun:"payment"`
  SystemBank      SystemBankRespPayment   `bun:"system_bank"`
  ImageSystemBank ImageSystemBankResp     `bun:"imagesystembank"`
  Shipment        ShipmentRespOrderDetail `bun:"shipment"`
  TotalAmount     int                     `json:"total_amount"`
  TotalPrice      float64                 `json:"total_price"`
  Status          string                  `json:"status"`
  Images          []int64                 `json:"images"`
  Created_at      int64                   `json:"created_at"`
  Updated_at      int64                   `json:"updated_at"`
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
// 	SystemBankID  int     `json:"system_bank_id" bun:"system_bank_id"`
// 	PaymentPrice  float64 `json:"payment_price" bun:"payment_price"`
// 	BankName      string  `json:"bank_name" bun:"bank_name"`
// 	AccountName   string  `json:"account_name" bun:"account_name"`
// 	AccountNumber string  `json:"account_number" bun:"account_number"`
// 	PaymentStatus string  `json:"payment_status" bun:"payment_status"`
