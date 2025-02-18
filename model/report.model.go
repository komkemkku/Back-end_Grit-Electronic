package model

import "github.com/uptrace/bun"

type Report struct {
	bun.BaseModel `bun:"table:report"`

	ID          int     `bun:",type:serial,autoincrement,pk"`
	ProductID   int     `bun:"product_id"`
	ProductName string  `bun:"product_name"`
	Date        int64   `bun:"date"`
	TotalPrice  float64 `bun:"total_price"`
	TotalAmount int     `bun:"total_amount"`
	Status      string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}