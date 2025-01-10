package requests

type ProductRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ProductIdRequest struct {
	ID int `uri:"id"`
}

type ProductCreateRequest struct {
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	Detail     string   `json:"detail"`
	Stock      int      `json:"stock"`
	Image      []string `json:"image"`
	Spec       string   `json:"spec"`
	CategoryID int      `json:"category_id"`
}

type ProductUpdateRequest struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	Detail     string   `json:"detail"`
	Stock      int      `json:"stock"`
	Image      []string `json:"image"`
	Spec       string   `json:"spec"`
	CategoryID int      `json:"category_id"`
}
