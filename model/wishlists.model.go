package model

import "github.com/uptrace/bun"

type Wishlists struct {
	bun.BaseModel `bun:"table:wishlists"`

	ID               int     `bun:",type:serial,autoincrement,pk"`
	UserID           int     `bun:"user_id"`
	ProductID        int     `bun:"product_id"`
	PricePerProduct  float32 `bun:"price_per_product"`
	AmountPerProduct int     `bun:"amount_per_product"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
