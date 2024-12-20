package model

import "github.com/uptrace/bun"

type Order_details struct {
	bun.BaseModel `bun:"table:order_detail"`

	Id          int64   `bun:",type:serial,autoincrement,pk"`
	Quantity    int64   `bun:"quantity"`
	Unit_price  float64 `bun:"unit_price"`
	Total_price float64 `bun:"total_price"`
	Order_id    int64   `bun:"bun:order_id"`
	Product_id  int64   `bun:"bun:product_id"`

	CreateUnixTimestamp
}
