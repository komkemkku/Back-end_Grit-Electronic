package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	TotalPrice  float64  `bun:"total_price"`
	TotalAmount int64  `bun:"total_amount"`
	Status      string `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
