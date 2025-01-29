package model

import "github.com/uptrace/bun"

type Carts struct {
	bun.BaseModel `bun:"table:carts"`

	ID               int     `bun:",type:serial,autoincrement,pk"`
	UserID           int     `bun:"user_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
