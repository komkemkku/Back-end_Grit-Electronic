package requests

type UserRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type UserIdRequest struct {
	ID int `uri:"id"`
}

type UserCreateRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserUpdateRequest struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
