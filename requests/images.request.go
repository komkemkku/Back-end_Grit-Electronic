package requests

type ImageRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ImageIdRequest struct {
	ID int `uri:"id"`
}

type ImageCreateRequest struct {
	ProductID    int  `json:"product_id"`
	ImageProduct string `json:"image_product"`
	ReviewID     int  `json:"review_id"`
	ImageReview  string `json:"image_review"`
}

type ImageUpdateRequest struct {
	ID           int  `json:"id"`
	ProductID    int  `json:"product_id"`
	ImageProduct string `json:"image_product"`
	ReviewID     int  `json:"review_id"`
	ImageReview  string `json:"image_review"`
}
