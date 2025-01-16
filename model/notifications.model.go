package model

import "github.com/uptrace/bun"

type Notifications struct {
	bun.BaseModel `bun:"table:notifications"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	UserID      int    `bun:"user_id"`
	AdminID     int    `bun:"admin_id"`
	Type        string `bun:"type"`
	Description string `bun:"description"`
	IsRead      bool   `bun:"is_read"`
	RefID       int    `bun:"ref_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
