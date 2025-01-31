package model

import "github.com/uptrace/bun"

type CartItem struct {
	bun.BaseModel `bun:"table:cart_items"`

	ID                 int     `bun:",type:serial,autoincrement,pk"`
	CartID             int     `bun:"cart_id"`
	ProductID          int     `bun:"product_id"`
	TotalProductAmount int     `bun:"total_product_amount"`
	// TotalProductPrice  float64 `bun:"total_product_price"`
	Status             string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
