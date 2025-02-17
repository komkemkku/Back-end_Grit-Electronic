package requests

type PaymentRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type PaymentIdRequest struct {
	ID int `uri:"id"`
}

type PaymentCreateRequest struct {
	Price        int    `json:"price"`
	SystemBankID int    `json:"system_bank_id"`
	Date         string `json:"date"`
}

type PaymentUpdateRequest struct {
	Id   int    `json:"id"`
	Date string `json:"date"`
}
