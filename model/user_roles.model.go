package model

import "github.com/uptrace/bun"

type UserRole struct {
	bun.BaseModel `bun:"table:user_roles"`

	UserID int64 `bun:"user_id"`
	RoleID int64 `bun:"role_id"`
}
