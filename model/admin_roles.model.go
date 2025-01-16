package model

import "github.com/uptrace/bun"

type AdminRoles struct {
	bun.BaseModel `bun:"table:admin_roles"`

	AdminID int `bun:"admin_id"`
	RoleID  int `bun:"role_id"`
}
