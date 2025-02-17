package response

type ImageBanner struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Banner     string `json:"banner"`
	Created_at int64  `json:"created_at"`
}

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

type ImageSystemBankResp struct {
	ID          int    `json:"id"`
	RefID       int    `json:"ref_id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type ImageCategoriesResp struct {
	ID          int    `json:"id"`
	RefID       int    `json:"ref_id"`
	Description string `json:"description"`
}

type ImagePaymentResp struct {
	ID          int    `json:"id"`
	RefID       int    `json:"ref_id"`
	Description string `json:"description"`
}

type ImageReviewResp struct {
	ID          int    `json:"id"`
	RefID       int    `json:"ref_id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
