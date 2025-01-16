package model

import "github.com/uptrace/bun"

type SystemBanks struct {
	bun.BaseModel `bun:"table:system_banks"`

	ID            int    `bun:",type:serial,autoincrement,pk"`
	BankName      string `bun:"bank_name"`
	AccountName   string `bun:"account_name"`
	AccountNumber string `bun:"account_number"`
	Description   string `bun:"description"`
	IsActive      bool   `bun:"is_active"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
