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
	ReviewText string `json:"review_text"`
	Rating      int64  `json:"rating"`
	ProductID  int64  `json:"product_id"`
	UserId      int64  `json:"user_id"`
}

type ReviewUpdateRequest struct {
	ID          int64  `json:"id"`
	ReviewText string `json:"review_text"`
	Rating      int64  `json:"rating"`
	ProductID  int64  `json:"product_id"`
	UserId      int64  `json:"user_id"`
}
