package response

type OrderResponses struct {
	ID                  int64   `json:"id"`
	UserID              int64   `json:"user_id"`
	Username            string  `json:"username"`
	UserFirstname       string  `json:"firstname"`
	UserLastname        string  `json:"lastname"`
	UserPhone           string  `json:"phone"`
	PaymentID           string  `json:"payment_id"`
	ShipmentID          int64   `json:"shipment_id"`
	ShipmentFirstname   string  `json:"shipment_firstname"`
	ShipmentLastname    string  `json:"shipment_lastname"`
	ShipmentAddress     string  `json:"shipment_address"`
	ShipmentZipCode     string  `json:"shipment_zip_code"`
	ShipmentSubDistrict string  `json:"shipment_sub_district"`
	ShipmentDistrict    string  `json:"shipment_district"`
	ShipmentProvince    string  `json:"shipment_province"`
	TotalPrice          float64 `json:"total_price"`
	TotalAmount         int     `json:"total_amount"`
	Status              string  `json:"status"`
	CreatedAt           int64   `json:"created_at"`
	UpdatedAt           int64   `json:"updated_at"`
}

type OrderRespOrderDetail struct {
	ID             int                     `json:"id"`
	User           UserRespOrderDetail     `bun:"user"`
	Products       []ProductInfo           `json:"products"` // แก้จาก []string เป็น []ProductInfo
	Payment        PaymentRespOrderDetail  `bun:"payment"`
	Shipment       ShipmentRespOrderDetail `bun:"shipment"`
	TotalAmount    int                     `json:"total_amount"`
	TotalPrice     float64                 `json:"total_price"`
	TrackingNumber string                  `json:"tracking_number"`
	Status         string                  `json:"status"`
	Created_at     int64                   `json:"created_at"`
	Updated_at     int64                   `json:"updated_at"`
}

// สร้าง struct เก็บข้อมูลสินค้า
type ProductInfo struct {
	ProductID          int64   `json:"product_id"`
	ProductName        string  `json:"product_name"`
	Price              float64 `json:"price"`
	TotalProductAmount int     `json:"total_product_amount"`
	Image              string  `json:"image"`
	IsReview           bool    `json:"is_review"`
}

type OrderRespReport struct {
	ID         int   `json:"id"`
	Created_at int64 `json:"created_at"`
}
