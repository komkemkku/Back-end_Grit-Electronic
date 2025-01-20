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
	Password  string `json:"_"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type UserUpdateRequest struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

