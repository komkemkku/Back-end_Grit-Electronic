package requests

type SystemBankRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type SystemBankIdRequest struct {
	ID int `uri:"id"`
}

type SystemBankCreateRequest struct {
	BankName        string `json:"bank_name"`
	AccountName     string `json:"account_name"`
	AccountNumber   string `json:"account_number"`
	Description     string `json:"description"`
	ImageSystemBank string `json:"image_system_bank"`
	IsActive        bool   `json:"is_active"`
}

type SystemBankUpdateRequest struct {
	ID              int    `json:"id"`
	BankName        string `json:"bank_name"`
	AccountName     string `json:"account_name"`
	AccountNumber   string `json:"account_number"`
	Description     string `json:"description"`
	ImageSystemBank string `json:"image_system_bank"`
	IsActive        bool   `json:"is_active"`
}
