package model

import "github.com/uptrace/bun"

type Role_Permissions struct {
	bun.BaseModel `bun:"table:role_permissions"`

	ID            int64 `bun:",type:serial,autoincrement,pk"`
	RoleID       int64 `bun:"bun:role_id"`
	PermissionID int64 `bun:"bun:permission_id"`

	CreateUnixTimestamp
}
