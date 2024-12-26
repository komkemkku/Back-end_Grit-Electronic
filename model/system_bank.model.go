package model

import "github.com/uptrace/bun"

type SystemBank struct {
	bun.BaseModel `bun:"table:system_bank"`

	ID             int64  `bun:",type:serial,autoincrement,pk"`
	BankName      string `bun:"bank_name"`
	AccountName   string `bun:"account_name"`
	AccountNumber string `bun:"account_number"`
	Image         string `bun:"image"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
