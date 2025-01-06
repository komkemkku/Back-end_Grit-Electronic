package model

import "github.com/uptrace/bun"

type Reports struct {
	bun.BaseModel `bun:"table:reports"`

	Id        int `bun:",type:serial,autoincrement,pk"`
	Sale      int `bun:"sale"`
	QtyReport int `bun:"qty_rep"`
	Profit    int `bun:"profit"`
	Cost      int `bun:"cost"`
	OrderID   int `bun:"bun:order_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
