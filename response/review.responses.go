package response

// type ReviewResponses struct {
// 	ID          int               `json:"id"`
// 	User        UserRespReview    `json:"user"`
// 	Product     ProductRespReview `json:"product"`
// 	Rating      int               `json:"rating"`
// 	TextReview  string            `json:"text_review"`
// 	ImageReview []string          `json:"image_review"`
// 	Created_at  int64             `json:"created_at"`
// 	Updated_at  int64             `json:"updated_at"`
// }

type ReviewProductResp struct {
	ID          int    `json:"id"`
	Rating      int    `json:"rating"`
}

type ReviewProductDetailResp struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

type ReviewResponses struct {
	ID          int64    `json:"id"`
	User        string   `json:"user"`    // จาก username
	Product     string   `json:"product"` // จาก product_name
	Rating      int      `json:"rating"`
	TextReview  string   `json:"text_review"` // จาก description
	ImageReview []string `json:"image_review"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
