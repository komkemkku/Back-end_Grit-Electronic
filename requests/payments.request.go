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
	Price         int  `json:"price"`
	Amount        int  `json:"amount"`
	Status        int `json:"status"`
	Slip          string `json:"slip"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}

type PaymentUpdateRequest struct {
	Id            int  `json:"id"`
	Price         int  `json:"price"`
	Amount        int  `json:"amount"`
	Status        int `json:"status"`
	Slip          string `json:"slip"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
}
