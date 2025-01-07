package model

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users"`

	ID         int  `bun:",type:serial,autoincrement,pk"`
	Username   string `bun:"username"`
	Password   string `bun:"password"`
	Email      string `bun:"email"`
	Phone      string `bun:"phone"`
	FirstName string `bun:"firstname"`
	LastName string `bun:"lastname"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
