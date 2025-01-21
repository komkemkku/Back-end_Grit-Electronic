package model

import "github.com/uptrace/bun"

type CartItem struct {
	bun.BaseModel `bun:"table:cart_items"`

	ID                  int     `bun:",type:serial,autoincrement,pk"`
	CartID              int     `bun:"cart_id"`
	ProductID           int     `bun:"product_id"`
	ProductName         string  `bun:"product_name"`
	ProductImageMain    string  `bun:"product_image_main"`
	TotalProductPrice   float32 `bun:"total_product_price"`
	TotalProductAamount int     `bun:"total_product_amount"`
	Status              string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
