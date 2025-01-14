package model

import "github.com/uptrace/bun"

type Order_details struct {
	bun.BaseModel `bun:"table:order_details"`

	ID         int   `bun:",type:serial,autoincrement,pk"`
	OrderID    int   `bun:"order_id"`
	ProductID  []int `bun:"product_id"`
	PaymentID  int   `bun:"payment_id"`
	ShipmentID int   `bun:"shipment_id"`

	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
