package requests

type ReviewRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ReviewIdRequest struct {
	Id int64 `uri:"id"`
}

type ReviewCreateRequest struct {
	Review_text string `json:"review_text"`
	Rating      int64  `json:"rating"`
	Product_id  int64  `json:"product_id"`
}

type ReviewUpdateRequest struct {
	Id          int64  `json:"id"`
	Review_text string `json:"review_text"`
	Rating      int64  `json:"rating"`
	Product_id  int64  `json:"product_id"`
}
