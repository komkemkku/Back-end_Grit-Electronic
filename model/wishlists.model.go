package model

import "github.com/uptrace/bun"

type Wishlists struct {
	bun.BaseModel `bun:"table:wishlists"`

	ID        int64 `bun:",type:serial,autoincrement,pk"`
	ProductID int64 `bun:"bun:product_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
