package model

import "github.com/uptrace/bun"

type Products struct {
	bun.BaseModel `bun:"table:products"`

	Id       int64  `bun:",type:serial,autoincrement,pk"`
	Name     string `bun:"name"`
	Price    int64  `bun:"price"`
	Details  string `bun:"details"`
	stock    int64  `bun:"stock"`
	Image string `bun:"image"`

	CreateUnixTimestamp
}
