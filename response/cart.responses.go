package response

type CartResponses struct {
	ID         int           `json:"id"`
	UserID     UserRespCart  `bun:"user"`
	CartItems  []CartItemRes `bun:"cart_items"`
	Created_at int64         `json:"created_at"`
	Updated_at int64         `json:"updated_at"`
}

