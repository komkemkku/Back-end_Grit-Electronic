package model

import "github.com/uptrace/bun"

type UserRole struct {
	bun.BaseModel `bun:"table:user_roles"`

	User_id int64 `bun:"user_id"`
	Role_id int64 `bun:"role_id"`
}
