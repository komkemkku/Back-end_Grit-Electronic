package model

import "github.com/uptrace/bun"

type Images struct {
	bun.BaseModel `bun:"table:images"`

	ID           int64  `bun:",type:serial,autoincrement,pk"`
	ProductID    int64  `bun:"bun:product_id"`
	ImageProduct string `bun:"image_product"`
	ReviewID     int64  `bun:"bun:review_id"`
	ImageReview  string `bun:"image_review"`

	CreateUnixTimestamp
}
