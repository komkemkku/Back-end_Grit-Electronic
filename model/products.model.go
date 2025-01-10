package model

import "github.com/uptrace/bun"

type Products struct {
	bun.BaseModel `bun:"table:products"`

	ID         int      `bun:",type:serial,autoincrement,pk"`
	Name       string   `bun:"name"`
	Price      float64  `bun:"price"`
	Detail     string   `bun:"detail"`
	Stock      int      `bun:"stock"`
	Image      []string `bun:"image"`
	Spec       string   `bun:"spec,notnull"`
	CategoryID int      `bun:"category_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
