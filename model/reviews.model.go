package model

import "github.com/uptrace/bun"

type Reviews struct {
	bun.BaseModel `bun:"table:reviews"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	ProductID   int64  `bun:"product_id"`
	UserID      int64  `bun:"user_id"`
	Description string `bun:"description"`
	Rating      int64  `bun:"rating"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
