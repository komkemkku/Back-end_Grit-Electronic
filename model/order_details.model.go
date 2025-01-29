package model

import "github.com/uptrace/bun"

type OrderDetail struct {
	bun.BaseModel `bun:"table:order_details"`

	ID                 int     `bun:",type:serial,autoincrement,pk"`
	OrderID            int     `bun:"order_id"`
	ProductName        string  `bun:"product_name"`
	TotalProductPrice  float64 `bun:"total_product_price"`
	TotalProductAmount int     `bun:"total_product_amount"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
