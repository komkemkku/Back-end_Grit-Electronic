package model

import "github.com/uptrace/bun"

type Carts struct {
	bun.BaseModel `bun:"table:carts"`

	ID        int `bun:",type:serial,autoincrement,pk"`
	Quantity  int `bun:"quantity"`
	ProductID int `bun:"bun:product_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
