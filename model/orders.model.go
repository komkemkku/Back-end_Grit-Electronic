package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`

	ID           int64  `bun:",type:serial,autoincrement,pk"`
	TotalPrice  int64  `bun:"total_price"`
	TotalAmount int64  `bun:"total_amount"`
	Status       string `bun:"status"`
	UserID      int64  `bun:"bun:user_id"`
	ShipmentID  int64  `bun:"bun:shipment_id"`
	PaymentID   int64  `bun:"bun:payment_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
