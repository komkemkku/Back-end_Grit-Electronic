package model

import "github.com/uptrace/bun"

type Reviews struct {
	bun.BaseModel `bun:"table:reviews"`

	Id          int64  `bun:",type:serial,autoincrement,pk"`
	Review_text string `bun:"review_text"`
	Rating      string `bun:"rating"`
	User_id     int64  `bun:"bun:user_id"`
	Product_id  int64  `bun:"bun:product_id"`

	CreateUnixTimestamp
}
