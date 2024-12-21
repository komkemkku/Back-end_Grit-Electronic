package requests

type CategoryRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CategoryIdRequest struct {
	Id int64 `uri:"id"`
}

type CategoryCreateRequest struct {
	Name string `json:"name"`
	Image string `json:"image"`
}

type CategoryUpdateRequest struct {
    Id     int64  `json:"id"`
    Name   string `json:"name"`
    Image string `json:"image"`
}