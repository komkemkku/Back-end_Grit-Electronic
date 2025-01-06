package model

import "github.com/uptrace/bun"

type Wishlists struct {
	bun.BaseModel `bun:"table:wishlists"`

	Id         int64 `bun:",type:serial,autoincrement,pk"`
	Product_id int64 `bun:"bun:product_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
