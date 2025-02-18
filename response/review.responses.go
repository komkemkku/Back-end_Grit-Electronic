package response

type ReviewProductResp struct {
	ID     int `json:"id"`
	Rating int `json:"rating"`
}

type ReviewProductDetailResp struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
}

type ReviewResponses struct {
	ID          int               `bun:"id" json:"id"`
	User        string            `bun:"user" json:"user"`
	Product     string            `bun:"product" json:"product"`
	Rating      int               `bun:"rating" json:"rating"`
	Description string            `bun:"description" json:"description"`
	CreatedAt   string            `bun:"created_at" json:"created_at"`
	UpdatedAt   string            `bun:"updated_at" json:"updated_at"`
}
