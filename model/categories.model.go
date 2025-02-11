package model

import "github.com/uptrace/bun"

type Categories struct {
	bun.BaseModel `bun:"table:categories"`

	ID       int    `bun:",type:serial,autoincrement,pk"`
	Name     string `bun:"name"`
	Image    string `bun:"image"`
	IsActive bool   `bun:"is_active"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
