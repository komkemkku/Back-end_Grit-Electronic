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
	AdminID int 
	Type   string `json:"type"`
	Banner string `json:"banner"`
}
