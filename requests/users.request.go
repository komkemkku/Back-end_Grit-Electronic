package requests

type UserRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type UserIdRequest struct {
	Id int64 `uri:"id"`
}

type UserCreateRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Bank_number string `json:"bank_number"`
	Role_id     string `json:"role_id"`
}

type UserUpdateRequest struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Bank_number string `json:"bank_number"`
	Role_id     string `json:"role_id"`
}
