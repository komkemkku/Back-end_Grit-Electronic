package requests

type ImageRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ImageIdRequest struct {
	Id int64 `uri:"id"`
}

type ImageCreateRequest struct {
	Product_id    int64  `json:"product_id"`
	Image_product string `json:"image_product"`
	Review_id     int64  `json:"review_id"`
	Image_review  string `json:"image_review"`
}

type ImageUpdateRequest struct {
	Id            int64  `json:"id"`
	Product_id    int64  `json:"product_id"`
	Image_product string `json:"image_product"`
	Review_id     int64  `json:"review_id"`
	Image_review  string `json:"image_review"`
}
