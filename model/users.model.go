package model

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users"`

	Id          int64  `bun:",type:serial,autoincrement,pk"`
	Username    string `bun:"username"`
	Password    string `bun:"password"`
	Email       string `bun:"email"`
	Phone       string `bun:"phone"`
	Bank_number string `bun:"bank_number"`

	CreateUnixTimestamp
}
