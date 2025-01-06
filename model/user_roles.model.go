package model

import "github.com/uptrace/bun"

type UserRole struct {
	bun.BaseModel `bun:"table:user_roles"`

	UserID int `bun:"user_id"`
	RoleID int `bun:"role_id"`
}
