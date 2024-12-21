package requests

type ProductRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ProductIdRequest struct {
	Id int64 `uri:"id"`
}

type ProductCreateRequest struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Detail  string `json:"detail"`
	Stock int64 `json:"stock"`
	Image string `json:"image"`
	Category_id int64 `json:"category_id"`
}

type ProductUpdateRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Detail string `json:"detail"`
	Stock int64 `json:"stock"`
	Image string `json:"image"`
	Category_id int64 `json:"category_id"`
}