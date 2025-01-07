package response

type CartResponses struct {
	ID         int64      `json:"id"`
	Quantity   int64      `json:"quantity"`
	TotalPrice float64    `json:"total_price"`
	Product    ProductResp `json:"product"`
	Created_at int64      `json:"created_at"`
	Updated_at int64      `json:"updated_at"`
}
