package model

import "github.com/uptrace/bun"

type Permissions struct {
	bun.BaseModel `bun:"table:permissions"`

	Id          int64  `bun:",type:serial,autoincrement,pk"`
	Group_name  string `bun:"group_name"`
	Name        string `bun:"name"`
	Description string `bun:"description"`

	CreateUnixTimestamp
}
