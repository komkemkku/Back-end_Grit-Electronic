package model

import "github.com/uptrace/bun"

type Permissions struct {
	bun.BaseModel `bun:"table:permissions"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	GroupName   string `bun:"group_name"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	IsActive    bool   `bun:"is_active"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
