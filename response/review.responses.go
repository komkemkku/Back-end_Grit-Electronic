package response

type ReviewResponses struct {
	ID          int               `json:"id"`
	User        UserRespReview    `json:"user"`
	Product     ProductRespReview `json:"product"`
	Description string            `json:"description"`
	Rating      int               `json:"rating"`
	Created_at  int64             `json:"created_at"`
	Updated_at  int64             `json:"updated_at"`
}

type ReviewProductResp struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
