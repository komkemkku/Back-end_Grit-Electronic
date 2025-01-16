package model

import "github.com/uptrace/bun"

type Images struct {
	bun.BaseModel `bun:"table:images"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	RefID       int    `bun:"ref_id"`
	Type        string `bun:"type"`
	Description string `bun:"description"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
