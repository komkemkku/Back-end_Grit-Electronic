package model

import "github.com/uptrace/bun"

type Reviews struct {
	bun.BaseModel `bun:"table:reviews"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	ReviewText string `bun:"review_text"`
	Rating      string `bun:"rating"`
	ProductID  int64  `bun:"bun:product_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
