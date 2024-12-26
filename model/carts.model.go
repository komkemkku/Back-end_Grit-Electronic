package model

import "github.com/uptrace/bun"

type Carts struct {
	bun.BaseModel `bun:"table:carts"`

	ID         int64 `bun:",type:serial,autoincrement,pk"`
	Quantity   int64 `bun:"quantity"`
	UserID    int64 `bun:"bun:user_id"`
	ProductID int64 `bun:"bun:product_id"`

	CreateUnixTimestamp
}
