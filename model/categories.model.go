package model

import "github.com/uptrace/bun"

type Categories struct {
	bun.BaseModel `bun:"table:categories"`

	ID       int    `bun:",type:serial,autoincrement,pk"`
	Name     string `bun:"name"`
	IsActive bool   `bun:"is_active"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
