package response

type PaymentResponses struct {
	ID            int     `json:"id"`
	Price         float64 `json:"price"`
	Amount        int     `json:"amount"`
	Slip          string  `json:"slip"`
	Status        int     `json:"status"`
	BankName      string  `json:"bank_name"`
	AccountName   string  `json:"account_name"`
	AccountNumber string  `json:"account_number"`
	Created_at    int64   `json:"created_at"`
	Updated_at    int64   `json:"updated_at"`
}

type PaymentRespOrderDetail struct {
	ID     int     `json:"id"`
	Price  float64 `json:"price"`
	Slip   string  `json:"slip"`
	Status int     `json:"status"`
}
