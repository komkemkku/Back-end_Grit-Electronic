package model

import "github.com/uptrace/bun"

type Products struct {
	bun.BaseModel `bun:"table:products"`

	ID          int64   `bun:",type:serial,autoincrement,pk"`
	Name        string  `bun:"name"`
	Price       float64 `bun:"price"`
	Detail      string  `bun:"detail"`
	Stock       int64   `bun:"stock"`
	Image       string  `bun:"image"`
	CategoryID int64   `bun:"category_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
