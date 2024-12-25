package requests

type SystemBankRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type SystemBankIdRequest struct {
	Id int64 `uri:"id"`
}

type SystemBankCreateRequest struct {
	Bank_name      string `json:"bank_name"`
	Account_name   string `json:"account_name"`
	Account_number string `json:"account_number"`
	Image          string `json:"image"`
}

type SystemBankUpdateRequest struct {
	Id             int64  `json:"id"`
	Bank_name      string `json:"bank_name"`
	Account_name   string `json:"account_name"`
	Account_number string `json:"account_number"`
	Image          string `json:"image"`
}
