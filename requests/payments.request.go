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
	AdminID       int    `json:"admin_id"`
	Price         int    `json:"price"`
	Amount        int    `json:"amount"`
	Status        int    `json:"status"`
	PaymentSlip   string `json:"payment_slip"`
	UpdatedBy     int    `json:"updated_by"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}

type PaymentUpdateRequest struct {
	Id            int    `json:"id"`
	AdminID       int    `json:"admin_id"`
	Price         int    `json:"price"`
	Amount        int    `json:"amount"`
	Status        int    `json:"status"`
	PaymentSlip   string `json:"payment_slip"`
	UpdatedBy     int    `json:"updated_by"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}
