package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`

	Id           int64  `bun:",type:serial,autoincrement,pk"`
	Total_Price  int64  `bun:"total_price"`
	Total_Amount int64  `bun:"total_amount"`
	Status       string `bun:"status"`
	User_id      int64  `bun:"bun:user_id"`
	Shipment_id  int64  `bun:"bun:shipment_id"`
	Payment_id   int64  `bun:"bun:payment_id"`

	CreateUnixTimestamp
}
