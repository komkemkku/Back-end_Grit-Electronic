package model

import "github.com/uptrace/bun"

type Shipments struct {
	bun.BaseModel `bun:"table:shipments"`

	ID          int    `bun:",type:serial,autoincrement,pk"`
	UserID      int    `bun:"user_id"`
	Firstname   string `bun:"firstname"`
	Lastname    string `bun:"lastname"`
	Address     string `bun:"address"`
	ZipCode     int    `bun:"zip_code"`
	SubDistrict string `bun:"sub_district"`
	District    string `bun:"district"`
	Province    string `bun:"province"`
	Status      string `bun:"status"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
