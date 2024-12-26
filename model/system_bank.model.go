package model

import "github.com/uptrace/bun"

type SystemBank struct {
	bun.BaseModel `bun:"table:system_bank"`

	Id             int64  `bun:",type:serial,autoincrement,pk"`
	Bank_name      string `bun:"bank_name"`
	Account_name   string `bun:"account_name"`
	Account_number string `bun:"account_number"`
	Image         string `bun:"image"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
