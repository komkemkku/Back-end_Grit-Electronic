package model

import "github.com/uptrace/bun"

type Products struct {
	bun.BaseModel `bun:"table:products"`

	ID          int     `bun:",type:serial,autoincrement,pk"`
	CategoryID  int     `bun:"category_id"`
	Name        string  `bun:"name"`
	Price       float64 `bun:"price"`
	Description string  `bun:"description"`
	Stock       int     `bun:"stock"`
	IsActive    bool    `bun:"is_active"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
