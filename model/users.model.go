package model

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	Username    string `bun:"username"`
	Password    string `bun:"password"`
	Email       string `bun:"email"`
	Phone       string `bun:"phone"`
	BankNumber string `bun:"bank_number"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}

type Admin struct {
	bun.BaseModel `bun:"table:users"`

	ID          int64  `bun:",type:serial,autoincrement,pk"`
	Username    string `bun:"username"`
	Password    string `bun:"password"`
	Email       string `bun:"email"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
}