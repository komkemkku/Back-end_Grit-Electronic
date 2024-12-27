package requests

type ImageRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ImageIdRequest struct {
	ID int64 `uri:"id"`
}

type ImageCreateRequest struct {
	ProductID    int64  `json:"product_id"`
	ImageProduct string `json:"image_product"`
	ReviewID     int64  `json:"review_id"`
	ImageReview  string `json:"image_review"`
}

type ImageUpdateRequest struct {
	ID           int64  `json:"id"`
	ProductID    int64  `json:"product_id"`
	ImageProduct string `json:"image_product"`
	ReviewID     int64  `json:"review_id"`
	ImageReview  string `json:"image_review"`
}
