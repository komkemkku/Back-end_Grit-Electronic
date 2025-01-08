package requests

type ReviewRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ReviewIdRequest struct {
	ID int64 `uri:"id"`
}

type ReviewCreateRequest struct {
	ReviewText  string `json:"review_text"`
	Rating      int    `json:"rating"`
	ProductID   int    `json:"product_id"`
	UserID      int64  `json:"user_id"`
	ImageReview []string `json:"image_review"`
}

type ReviewUpdateRequest struct {
	ID          int64  `json:"id"`
	ReviewText  string `json:"review_text"`
	Rating      int    `json:"rating"`
	ProductID   int64  `json:"product_id"`
	UserID      int64  `json:"user_id"`
	ImageReview []string `json:"image_review"`
}
