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
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
}

type ProductUpdateRequest struct {
	Id          int     `json:"id"`
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
}
