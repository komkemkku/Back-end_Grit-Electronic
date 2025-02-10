package response

type ProductResponses struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Price       float64             `json:"price"`
	Stock       int                 `json:"stock"`
	Description string              `json:"description"`
	Image       ImageProductResp    `bun:"image"`
	Category    CategoryProductResp `json:"category"`
	Review      []ReviewProductResp `bun:"reviews"`
	IsActive    bool                `json:"is_active"`
	Created_at  int64               `json:"created_at"`
	Updated_at  int64               `json:"updated_at"`
	Deleted_at  int64               `json:"deleted_at"`
}

type ProductDetailResponses struct {
	ID          int                       `json:"id"`
	Name        string                    `json:"name"`
	Price       float64                   `json:"price"`
	Stock       int                       `json:"stock"`
	Description string                    `json:"description"`
	Image       ImageProductDetailResp    `bun:"image"`
	Category    CategoryProductResp       `json:"category"`
	Review      []ReviewProductDetailResp `bun:"reviews"`
	IsActive    bool                      `json:"is_active"`
	Created_at  int64                     `json:"created_at"`
	Updated_at  int64                     `json:"updated_at"`
	Deleted_at  int64                     `json:"deleted_at"`
}

type ProductRespCart struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductRespCartItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductRespReview struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// type ProductDetail struct {
//     ProductID   int    `json:"product_id"`
//     ProductName string `json:"product_name"`
// }