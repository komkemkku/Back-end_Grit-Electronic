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




