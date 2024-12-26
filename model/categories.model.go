package model

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID    int64  `bun:",type:serial,autoincrement,pk"`
	Name  string `bun:"name"`
	Image string `bun:"image"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
