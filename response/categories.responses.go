package response

type CategoryResponses struct {
	ID              int                 `json:"id"`
	Name            string              `json:"name"`
	IsActive        bool                `json:"is_active"`
	ImageCategories ImageCategoriesResp `bun:"image"`
	Created_at      int64               `json:"created_at"`
	Updated_at      int64               `json:"updated_at"`
}

type CategoryProductResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
