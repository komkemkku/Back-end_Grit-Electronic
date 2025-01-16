package model

import "github.com/uptrace/bun"

type Admins struct {
	bun.BaseModel `bun:"table:admins"`

	ID       int    `bun:",type:serial,autoincrement,pk"`
	RoleID   int    `bun:"role_id"`
	Name     string `bun:"name"`
	Email    string `bun:"email"`
	Password string `bun:"password"`
	IsActive bool   `bun:"is_active"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
