package model

import "github.com/uptrace/bun"

type Payments struct {
	bun.BaseModel `bun:"table:payments"`

	ID            int     `bun:",type:serial,autoincrement,pk"`
	SystemBankID  int     `bun:"system_bank_id"`
	// Price         float64 `bun:"price"`
	// UpdatedBy     int     `bun:"updated_by"`
	Date          string  `bun:"date"`
	// BankName      string  `bun:"bank_name"`
	// AccountName   string  `bun:"account_name"`
	// AccountNumber string  `bun:"account_number"`
	// Status        string  `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}
