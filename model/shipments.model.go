package model

import "github.com/uptrace/bun"

type Shipments struct {
	bun.BaseModel `bun:"table:shipments"`

	Id           int64  `bun:",type:serial,autoincrement,pk"`
	Address      string `bun:"address"`
	Zip_code     string `bun:"zip_code"`
	Sub_district string `bun:"sub_district"`
	District     string `bun:"district"`
	Province     string `bun:"province"`
	User_id      int64  `bun:"bun:user_id"`
	Order_id     int64  `bun:"bun:order_id"`

	CreateUnixTimestamp
}
