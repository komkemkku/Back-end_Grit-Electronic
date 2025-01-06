package model

import "github.com/uptrace/bun"

type Reviews struct {
	bun.BaseModel `bun:"table:reviews"`

	ID         int    `bun:",type:serial,autoincrement,pk"`
	ReviewText string `bun:"text"`
	Rating     int    `bun:"rating"`
	ProductID  int    `bun:"bun:product_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
