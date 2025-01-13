package model

import "github.com/uptrace/bun"

type Orders struct {
	bun.BaseModel `bun:"table:orders"`

	ID          int     `bun:",type:serial,autoincrement,pk"`
	TotalPrice  float64 `bun:"total_price"`
	TotalAmount int     `bun:"total_amount"`
	Status      int     `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
