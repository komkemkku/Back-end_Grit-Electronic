package response

type ProductResponses struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	// Image       ImageProductResp    `bun:"image"`
	Category   CategoryProductResp `json:"category"`
	Review     []ReviewProductResp `bun:"reviews"`
	IsActive   bool                `json:"is_active"`
	Created_at int64               `json:"created_at"`
	Updated_at int64               `json:"updated_at"`
	Deleted_at int64               `json:"deleted_at"`
}

type ProductDetailResponses struct {
	ID          int                       `json:"id"`
	Name        string                    `json:"name"`
	Price       float64                   `json:"price"`
	Stock       int                       `json:"stock"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	Category    CategoryProductResp       `json:"category"`
	Review      []ReviewProductDetailResp `bun:"reviews"`
	IsActive    bool                      `json:"is_active"`
	IsFavorite  bool                      `json:"is_favorite"`
	Created_at  int64                     `json:"created_at"`
	Updated_at  int64                     `json:"updated_at"`
	Deleted_at  int64                     `json:"deleted_at"`
}

type ProductRespCart struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type ProductRespCartItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ProductRespReview struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductRespOrder struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductRespWish struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}
