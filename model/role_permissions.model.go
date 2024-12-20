package model

import "github.com/uptrace/bun"

type Role_Permissions struct {
	bun.BaseModel `bun:"table:role_permissions"`

	Id            int64 `bun:",type:serial,autoincrement,pk"`
	Role_id       int64 `bun:"bun:role_id"`
	Permission_id int64 `bun:"bun:permission_id"`

	CreateUnixTimestamp
}
