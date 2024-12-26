package model

import "github.com/uptrace/bun"

type Roles struct {
	bun.BaseModel `bun:"table:roles"`

	ID    int64  `bun:",type:serial,autoincrement,pk"`
	Name  string `bun:"name"`

	CreateUnixTimestamp
}