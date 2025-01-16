package model

import "github.com/uptrace/bun"

type Users struct {
	bun.BaseModel `bun:"table:users"`

	ID        int    `bun:",type:serial,autoincrement,pk"`
	Username  string `bun:"username"`
	Password  string `bun:"password"`
	Email     string `bun:"email"`
	Phone     string `bun:"phone"`
	FirstName string `bun:"firstname"`
	LastName  string `bun:"lastname"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}

// type Users struct {
// 	ID        int64  `json:"id"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Username  string `json:"username"`
// 	Password  string `json:"-"` // บอกให้ข้ามฟิลด์นี้ใน JSON Response
// 	Email     string `json:"email"`
// 	Phone     string `json:"phone"`
// 	CreatedAt int64  `json:"created_at"`
// 	UpdatedAt int64  `json:"updated_at"`

// 	CreateUnixTimestamp
// 	UpdateUnixTimestamp
// 	SoftDelete
// }
