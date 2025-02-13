package response

type WishlistResponses struct {
	ID        int              `json:"id"`
	User      UserRespWistlist `json:"user"`
	Product   ProductRespWish  `json:"product"`
	CreatedAt int64            `json:"created_at"`
	UpdatedAt int64            `json:"updated_at"`
}
