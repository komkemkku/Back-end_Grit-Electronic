package response

type UserResponses struct {
	ID        int           `json:"id"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Created_at int64         `json:"create_at"`
	Updated_at string        `json:"updated_at"`
}

type UserRespReview struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

type UserRespWistlist struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserRespCart struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}