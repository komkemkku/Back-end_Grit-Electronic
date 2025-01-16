package response

type AdminResponses struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Role       RoleResponses `json:"role"`
	Created_at int64         `json:"created_at"`
	Updated_at int64         `json:"updated_at"`
}

type AdminLogResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AdminPaymentResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
