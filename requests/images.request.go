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
	RefID       int    `json:"ref_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
