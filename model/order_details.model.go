package model

import "github.com/uptrace/bun"

type Order_details struct {
	bun.BaseModel `bun:"table:order_detail"`

	ID          int64   `bun:",type:serial,autoincrement,pk"`
	Quantity    int64   `bun:"quantity"`
	UnitPrice  float64 `bun:"unit_price"`
	TotalPrice float64 `bun:"total_price"`
	OrderID    int64   `bun:"bun:order_id"`
	ProductID  int64   `bun:"bun:product_id"`

	CreateUnixTimestamp
}
