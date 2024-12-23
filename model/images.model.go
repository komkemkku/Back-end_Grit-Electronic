package model

import "github.com/uptrace/bun"

type Images struct {
	bun.BaseModel `bun:"table:images"`

	Id            int64  `bun:",type:serial,autoincrement,pk"`
	Product_id    int64  `bun:"bun:product_id"`
	Image_product string `bun:"image_product"`
	Review_id     int64  `bun:"bun:review_id"`
	Image_review  string `bun:"image_review"`

	CreateUnixTimestamp
}
