package response

type UserResponses struct {
	ID        int           `json:"id"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Create_at int64         `json:"create_at"`
	UpdatedAt string        `json:"updated_at"`
}

type UserRespReview struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

type UserRespWistlist struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
