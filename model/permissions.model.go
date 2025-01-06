package model

import "github.com/uptrace/bun"

type Permissions struct {
	bun.BaseModel `bun:"table:permissions"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	GroupName   string `bun:"group_name"`
	Name        string `bun:"name"`
	Description string `bun:"description"`

	CreateUnixTimestamp
}
