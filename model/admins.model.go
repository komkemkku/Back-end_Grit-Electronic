package model

import "github.com/uptrace/bun"

type Admins struct {
	bun.BaseModel `bun:"table:admin"`

	Id      int64 `bun:",type:serial,autoincrement,pk"`
	User_id int64 `bun:"bun:user_id"`
	Role_id int64 `bun:"bun:role_id"`

	CreateUnixTimestamp
}
