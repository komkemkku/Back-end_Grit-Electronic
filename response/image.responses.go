package response

type ImageProductDetailResp struct {
	ID          int    `json:"id"`
	RefID       int    `json:"ref_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ImageProductResp struct {
	ID          int    `json:"id"`
	RefID       int    `json:"ref_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
