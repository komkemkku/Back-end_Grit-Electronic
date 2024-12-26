package model

import "github.com/uptrace/bun"

type Shipments struct {
	bun.BaseModel `bun:"table:shipments"`

	ID           int64  `bun:",type:serial,autoincrement,pk"`
	Address      string `bun:"address"`
	ZipCode     string `bun:"zip_code"`
	SubDistrict string `bun:"sub_district"`
	District     string `bun:"district"`
	Province     string `bun:"province"`
	UserID      int64  `bun:"bun:user_id"`
	OrderID     int64  `bun:"bun:order_id"`

	CreateUnixTimestamp
}
