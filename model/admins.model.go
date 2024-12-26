package model

import "github.com/uptrace/bun"

type Admins struct {
	bun.BaseModel `bun:"table:admin"`

	ID      int64 `bun:",type:serial,autoincrement,pk"`
	UserID int64 `bun:"bun:user_id"`
	RoleID int64 `bun:"bun:role_id"`

	CreateUnixTimestamp
}
