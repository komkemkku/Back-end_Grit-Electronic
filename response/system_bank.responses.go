package response

type SystemBankResponses struct {
	ID             int    `json:"id"`
	Bank_name      string `json:"bank_name"`
	Account_name   string `json:"account_name"`
	Account_number string `json:"account_number"`
	Description    string `json:"description"`
	Image          string `json:"image"`
	// ImageSystemBank ImageSystemBankResp `bun:"image"`
	IsActive   bool  `json:"is_active"`
	Created_at int64 `json:"created_at"`
	Updated_at int64 `json:"updated_at"`
}

type SystemBankRespPayment struct {
	ID            int64  `json:"id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Description   string `json:"description"`
	Image         string `json:"image"`
}
