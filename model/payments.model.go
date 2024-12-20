package model

import "github.com/uptrace/bun"

type Payments struct {
	bun.BaseModel `bun:"table:payments"`

	Id       int64   `bun:",type:serial,autoincrement,pk"`
	Price    float64 `bun:"price"`
	Amount   int64   `bun:"amount"`
	Slip     string  `bun:"slip"`
	User_id  int64   `bun:"bun:user_id"`
	Order_id int64   `bun:"bun:order_id"`

	CreateUnixTimestamp
}
