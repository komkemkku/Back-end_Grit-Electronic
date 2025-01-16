package model

import "github.com/uptrace/bun"

type OrderDetails struct {
	bun.BaseModel `bun:"table:order_details"`

	ID               int     `bun:",type:serial,autoincrement,pk"`
	OrderID          int     `bun:"order_id"`
	ProductID        int     `bun:"product_id"`
	ProductName      string  `bun:"product_name"`
	ProductImageMain string  `bun:"product_image_main"`
	PricePerProduct  float32 `bun:"price_per_product"`
	AmountPerProduct int     `bun:"amount_per_product"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
