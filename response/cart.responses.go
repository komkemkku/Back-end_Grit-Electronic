package response

type CartResponses struct {
	ID         int             `json:"id"`
	Quantity   int             `json:"quantity"`
	TotalPrice float64         `json:"total_price"`
	Product    ProductRespCart `json:"product"`
	Created_at int64           `json:"created_at"`
	Updated_at int64           `json:"updated_at"`
}
