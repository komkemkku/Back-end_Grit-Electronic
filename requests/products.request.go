package requests

type ProductRequest struct {
	Page       int64  `form:"page"`
	Size       int64  `form:"size"`
	Search     string `form:"search"`
	CategoryID int    `form:"category"`
}

type ProductIdRequest struct {
	ID int64 `uri:"id"`
}

type ProductUserIDRequest struct {
	UserID int64 `form:"user_id"`
}

type ProductCreateRequest struct {
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	IsActive    bool    `json:"is_active"`
}

type ProductUpdateRequest struct {
	Id          int     `json:"id"`
	CategoryID  int     `json:"category_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Image       string  `json:"image"`
	IsActive    bool    `json:"is_active"`
}
