package model

import "github.com/uptrace/bun"

type Reviews struct {
	bun.BaseModel `bun:"table:reviews"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	ProductID   int    `bun:"product_id"`
	UserID      int    `bun:"user_id"`
	Description string `bun:"description"`
	Rating      int    `bun:"rating"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
