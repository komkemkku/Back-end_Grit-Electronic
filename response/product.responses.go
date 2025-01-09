package response

type ProductResponses struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	Price      float64             `json:"price"`
	Detail     string              `json:"detail"`
	Spec       string              `json:"spec"`
	Stock      int                 `json:"stock"`
	Image      string              `json:"image"`
	Category   CategoryProductResp `json:"category"`
	Review     ReviewProductResp   `json:"review"`
	Created_at int64               `json:"created_at"`
	Updated_at int64               `json:"updated_at"`
}

type ProductRespCart struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ProductRespReview struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
