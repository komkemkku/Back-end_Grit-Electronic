package model

import "github.com/uptrace/bun"

type Reports struct {
	bun.BaseModel `bun:"table:reports"`

	Id       int64 `bun:",type:serial,autoincrement,pk"`
	Sale     int64 `bun:"sale"`
	QtyReport   int64 `bun:"qty_re"`
	Profit   int64 `bun:"profit"`
	Cost     int64 `bun:"cost"`
	OrderID int64 `bun:"bun:order_id"`

	CreateUnixTimestamp
}
