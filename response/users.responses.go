package response

type UserResponses struct {
	Id          int           `json:"id"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Phone       string        `json:"phone"`
	Bank_number string        `json:"bank_number"`
	Create_at   int64         `json:"create_at"`
	Role        RoleResponses `json:"role"`
}
