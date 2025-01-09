package response

type ReviewResponses struct {
	ID          int               `json:"id"`
	Username    UserRespReview    `json:"user"`
	Product     ProductRespReview `json:"product"`
	Rating      int               `json:"rating"`
	TextReview  string            `json:"text_review"`
	ImageReview []string          `json:"image_review"`
	Created_at  int64             `json:"created_at"`
	Updated_at  int64             `json:"updated_at"`
}

type ReviewProductResp struct {
	Username    string   `json:"username"`
	Rating      int      `json:"rating"`
	TextReview  string   `json:"text_review"`
}
