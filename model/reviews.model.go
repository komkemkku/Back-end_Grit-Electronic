package model

import "github.com/uptrace/bun"

type Reviews struct {
	bun.BaseModel `bun:"table:reviews"`

	ID          int64    `bun:",type:serial,autoincrement,pk"`
	TextReview  string   `bun:"text_review"`
	Rating      int64    `bun:"rating"`
	ProductID   int64    `bun:"product_id"`
	UserID      int64    `bun:"user_id"`
	ImageReview []string `bun:"image_review,type:jsonb"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
