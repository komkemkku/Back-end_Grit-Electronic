package response

type OrderResponses struct {
	ID         int               `json:"id"`
	User       UserRespCart      `bun:"user"`
	Payment    PaymentOrderResp  `bun:"payment"`
	Shipment   ShipmentRespOrder `bun:"shipment"`
	Cart       CartResponses     `bun:"cart"`
	Status     string            `json:"status"`
	Created_at int64             `json:"created_at"`
	Updated_at int64             `json:"updated_at"`
}

type OrderRespOrderDetail struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}
