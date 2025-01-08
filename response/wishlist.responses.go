package response

type WishlistResponses struct {
	ID         int               `json:"id"`
	Product    ProductRespReview `json:"product"`
	Created_at int64             `json:"created_at"`
	Updated_at int64             `json:"updated_at"`
}
