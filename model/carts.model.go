package model

import "github.com/uptrace/bun"

type Carts struct {
	bun.BaseModel `bun:"table:carts"`

	ID               int     `bun:",type:serial,autoincrement,pk"`
	UserID           int     `bun:"user_id"`
	ProductID        int     `bun:"product_id"`
	PricePerProduct  float32 `bun:"price_per_product"`
	AmountPerProduct float32 `bun:"amount_per_product"`
	TotalCartAmount  int     `bun:"total_cart_amount"`
	TotalCartPrice   float32 `bun:"total_cart_price"`
	Status           string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
