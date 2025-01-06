package model

import "github.com/uptrace/bun"

type Order_details struct {
	bun.BaseModel `bun:"table:order_detail"`

	ID        int64   `bun:",type:serial,autoincrement,pk"`
	Quantity  int64   `bun:"quantity"`
	Price     float64 `bun:"unit_price"`
	OrderID   int64   `bun:"order_id"`
	ProductID int64   `bun:"product_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
