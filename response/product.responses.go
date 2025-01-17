package response

type ProductResponses struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Price       float64             `json:"price"`
	Stock       int                 `json:"stock"`
	Description string              `json:"description"`
	Category    CategoryProductResp `json:"category"`
	// Reviews    []ReviewProductResp `json:"reviews"`
	IsActive   bool  `json:"is_active"`
	Created_at int64 `json:"created_at"`
	Updated_at int64 `json:"updated_at"`
}

type ProductRespCart struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price"`
}

type ProductRespReview struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductRespOrderDetail struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
