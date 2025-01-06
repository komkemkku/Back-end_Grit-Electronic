package model

import "github.com/uptrace/bun"

type Images struct {
	bun.BaseModel `bun:"table:images"`

	ID           int  `bun:",type:serial,autoincrement,pk"`
	ProductID    int  `bun:"bun:product_id"`
	ImageProduct string `bun:"product_img"`
	ReviewID     int  `bun:"bun:review_id"`
	ImageReview  string `bun:"review_img"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
