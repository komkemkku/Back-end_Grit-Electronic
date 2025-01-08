package response

type UserResponses struct {
	ID        int           `json:"id"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Create_at int64         `json:"create_at"`
	Role      RoleResponses `json:"role"`
}

type UserRespReview struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
