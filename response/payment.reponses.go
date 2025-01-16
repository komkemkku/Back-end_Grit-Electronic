package response

type PaymentResponses struct {
	ID            int              `json:"id"`
	Price         float64          `json:"price"`
	PaymentSlip   string           `json:"payment_slip"`
	Status        int              `json:"status"`
	UpdatedBy     int              `json:"updated_by"`
	AdminID       AdminPaymentResp `json:"admin"`
	BankName      string           `json:"bank_name"`
	AccountName   string           `json:"account_name"`
	AccountNumber string           `json:"account_number"`
	Created_at    int64            `json:"created_at"`
	Updated_at    int64            `json:"updated_at"`
}

type PaymentRespOrderDetail struct {
	ID          int     `json:"id"`
	Price       float64 `json:"price"`
	PaymentSlip string  `json:"payment_slip"`
	Status      int     `json:"status"`
}
