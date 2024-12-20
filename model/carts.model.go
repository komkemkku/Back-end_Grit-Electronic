package model

import "github.com/uptrace/bun"

type Carts struct {
	bun.BaseModel `bun:"table:carts"`

	Id         int64 `bun:",type:serial,autoincrement,pk"`
	Quantity   int64 `bun:"quantity"`
	User_id    int64 `bun:"bun:user_id"`
	Product_id int64 `bun:"bun:product_id"`

	CreateUnixTimestamp
}
