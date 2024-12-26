package response

type SystemBankResponses struct {
	Id             int64  `json:"id"`
	Bank_name      string `json:"bank_name"`
	Account_name   string `json:"account_name"`
	Account_number string `json:"account_number"`
	Image          string `json:"image"`
	Created_at     int64  `json:"created_at"`
	Updated_at     int64  `json:"updated_at"`
}
