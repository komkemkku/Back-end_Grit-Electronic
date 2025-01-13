package model

import "github.com/uptrace/bun"

type AdminLogs struct {
	bun.BaseModel `bun:"table_logs"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	AdminID     int    `bun:"admin_id"`
	Action      string `bun:"action"`
	Description string `bun:"description"`
	CreateUnixTimestamp
}
