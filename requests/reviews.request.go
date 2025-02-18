package requests

type ReviewRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ReviewIdRequest struct {
	ID int `uri:"id"`
}

type ReviewCreateRequest struct {
	ProductID   int      `json:"product_id"`
	UserID      int      `json:"user_id"`
	Description string   `json:"description"`
	Rating      int      `json:"rating"`
}

type ReviewUpdateRequest struct {
	ID          int      `json:"id"`
	ProductID   int      `json:"product_id"`
	UserID      int      `json:"user_id"`
	Description string   `json:"description"`
	Rating      int      `json:"rating"`
}
