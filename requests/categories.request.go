package requests

type CategoryRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type CategoryIdRequest struct {
	ID int `uri:"id"`
}

type CategoryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	IsActive    bool   `json:"is_active"`
	// ImageCategories string `json:"image_categories"`
}

type CategoryUpdateRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	IsActive bool   `json:"is_active"`
	// ImageCategories string `json:"image_categories"`
}
