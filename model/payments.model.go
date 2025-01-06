package model

import "github.com/uptrace/bun"

type Payments struct {
	bun.BaseModel `bun:"table:payments"`

	ID            int     `bun:",type:serial,autoincrement,pk"`
	Price         float64 `bun:"price"`
	Amount        int     `bun:"amount"`
	Slip          string  `bun:"slip"`
	Status        int     `bun:"status"`
	BankName      string  `bun:"bank_name"`
	AccountName   string  `bun:"account_name"`
	AccountNumber string  `bun:"account_number"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
