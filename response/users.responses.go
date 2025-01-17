package response

type UserResponses struct {
	ID        int           `json:"id"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Username  string        `json:"username"`
	Password  string        `json:"-"` // Hide password in response
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Create_at int64         `json:"create_at"`
	UpdatedAt string        `json:"updated_at"`
	Role      RoleResponses `json:"role"`
}

type UserRespReview struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

type UserRespWistlist struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
