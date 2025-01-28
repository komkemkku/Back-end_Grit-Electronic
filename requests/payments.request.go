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
	Price         int    `json:"price"`
	SystemBankID  int    `json:"system_bank_id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	PaymentSlip   string `json:"payment_slip"`
	Status        string `json:"status"`
}

type PaymentUpdateRequest struct {
	Id            int    `json:"id"`
	Price         int    `json:"price"`
	UpdatedBy     int    `json:"updated_by"`
	SystemBankID  int    `json:"system_bank_id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	PaymentSlip   string `json:"payment_slip"`
	Status        string `json:"status"`
}
