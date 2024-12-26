package model

import "github.com/uptrace/bun"

type Payments struct {
	bun.BaseModel `bun:"table:payments"`

	ID       int64   `bun:",type:serial,autoincrement,pk"`
	Price    float64 `bun:"price"`
	Amount   int64   `bun:"amount"`
	Slip     string  `bun:"slip"`
	UserID  int64   `bun:"bun:user_id"`
	OrderID int64   `bun:"bun:order_id"`

	CreateUnixTimestamp
}
