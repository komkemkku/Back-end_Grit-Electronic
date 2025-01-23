package model


import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`

	ID          int     `bun:",type:serial,autoincrement,pk"`
	UserID      int     `bun:"user_id"`
	PaymentID   int     `bun:"payment_id"`
	ShipmentID  int     `bun:"shipment_id"`
	CartID      int     `bun:"cart_id"`
	Status      string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}

// type Orders struct {
// 	bun.BaseModel `bun:"table:orders"`

// 	ID         int    `bun:",pk,autoincrement"`
// 	UserID     int    `bun:"user_id"`
// 	PaymentID  int    `bun:"payment_id"`
// 	ShipmentID int    `bun:"shipment_id"`
// 	CartID     int    `bun:"cart_id"`
// 	Status     string `bun:"status"` // หากเก็บเป็น string

// 	CreatedAt int64 `bun:"created_at"` // หรือ time.Time ขึ้นอยู่กับการใช้งาน
// 	UpdatedAt int64 `bun:"updated_at"` // หรือ time.Time
// }

// ตัวอย่างฟังก์ชัน SetCreatedNow/SetUpdateNow
// func (o *Orders) SetCreatedNow() {
// 	o.CreatedAt = time.Now().Unix()
// }

// func (o *Orders) SetUpdateNow() {
// 	o.UpdatedAt = time.Now().Unix()
// }

// Orders struct - โครงสร้างของคำสั่งซื้อ
// type Orders struct {
// 	ID         int64     `json:"id"`
// 	UserID     int64     `json:"user_id"`
// 	PaymentID  int64     `json:"payment_id"`
// 	ShipmentID int64     `json:"shipment_id"`
// 	CartID     int64     `json:"cart_id"`
// 	Status     string    `json:"status"`

// 	CreateUnixTimestamp
// 	UpdateUnixTimestamp
// }

// // SetCreatedNow - ตั้งค่า CreatedAt ให้เป็นเวลาปัจจุบัน
// func (o *Orders) SetCreatedNow() {
// 	o.CreatedAt = time.Now()
// }

// // SetUpdatedNow - ตั้งค่า UpdatedAt ให้เป็นเวลาปัจจุบัน
// func (o *Orders) SetUpdatedNow() {
// 	o.UpdatedAt = time.Now()
// }

// // SetUpdateNow - ตั้งค่า UpdatedAt ให้เป็นเวลาปัจจุบัน
// func (o *Orders) SetUpdateNow() {
// 	o.UpdatedAt = time.Now()
// }
