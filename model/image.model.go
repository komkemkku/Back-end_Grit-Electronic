package model

import "github.com/uptrace/bun"

type Images struct {
	bun.BaseModel `bun:"table:images"`

	ID     int    `bun:",type:serial,autoincrement,pk"`
	Type   string `bun:"type"`
	Banner string `bun:"banner"`

	CreateUnixTimestamp
}
