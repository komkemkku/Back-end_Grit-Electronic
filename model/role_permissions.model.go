package model

import "github.com/uptrace/bun"

type RolePermissions struct {
	bun.BaseModel `bun:"table:role_permissions"`

	RoleID       int `bun:"bun:role_id"`
	PermissionID int `bun:"bun:permission_id"`

}
