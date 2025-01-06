package model

import "github.com/uptrace/bun"

type Admins struct {
	bun.BaseModel `bun:"table:admins"`

	ID       int  `bun:",type:serial,autoincrement,pk"`
	Name     string `bun:"name"`
	Password string `bun:"password"`
	Email    string `bun:"email"`
	RoleID   int  `bun:"role_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
