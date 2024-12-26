package requests

type SystemBankRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type SystemBankIdRequest struct {
	ID int64 `uri:"id"`
}

type SystemBankCreateRequest struct {
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Image          string `json:"image"`
}

type SystemBankUpdateRequest struct {
	Id             int64  `json:"id"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	Image          string `json:"image"`
}
