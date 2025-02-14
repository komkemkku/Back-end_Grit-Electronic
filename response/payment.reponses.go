package response

type PaymentResponses struct {
	ID              int                   `json:"id"`
	UpdatedBy       int                   `json:"updated_by"`
	SystemBank      SystemBankRespPayment `bun:"systembank"`
	ImageSystemBank ImageSystemBankResp   `bun:"imagesystembank"`
	Price           float64               `json:"price"`
	Status          string                `json:"status"`
	Image           ImagePaymentResp      `bun:"image"`
	BankName        string                `json:"bank_name"`
	AccountName     string                `json:"account_name"`
	AccountNumber   string                `json:"account_number"`
	Created_at      int64                 `json:"created_at"`
	Updated_at      int64                 `json:"updated_at"`
}

type PaymentUserResp struct {
	ID         int    `json:"id"`
	SystemBank int    `json:"system_bank"`
	Date       string `json:"date"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
}

type PaymentOrderResp struct {
	ID            int              `json:"id"`
	UpdatedBy     int              `json:"updated_by"`
	Price         float64          `json:"price"`
	Status        string           `json:"status"`
	Image         ImagePaymentResp `bun:"image"`
	BankName      string           `json:"bank_name"`
	AccountName   string           `json:"account_name"`
	AccountNumber string           `json:"account_number"`
	Description   string           `json:"description"`
}

// type PaymentRespOrderDetail struct {
// 	ID            int     `json:"id"`
// 	Price         float64 `json:"price"`
// 	BankName      string  `json:"bank_name"`
// 	AccountName   string  `json:"account_name"`
// 	AccountNumber string  `json:"account_number"`
// 	Status        string  `json:"status"`
// }

type PaymentRespOrderDetail struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}
