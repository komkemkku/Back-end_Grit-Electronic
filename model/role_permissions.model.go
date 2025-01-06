package model

import "github.com/uptrace/bun"

type Role_Permissions struct {
	bun.BaseModel `bun:"table:role_permissions"`

	ID           int `bun:",type:serial,autoincrement,pk"`
	RoleID       int `bun:"bun:role_id"`
	PermissionID int `bun:"bun:permission_id"`

	CreateUnixTimestamp
}
