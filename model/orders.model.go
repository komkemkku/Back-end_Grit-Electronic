package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`

	ID           int     `bun:",type:serial,autoincrement,pk"`
	UserID       int     `bun:"user_id"`
	PaymentID    int     `bun:"payment_id"`
	ShipmentID   int     `bun:"shipment_id"`
	Total_price  float64 `bun:"total_price"`
	Total_amount int     `bun:"total_amount"`
	Status       string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
