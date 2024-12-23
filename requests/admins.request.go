package requests

type AdminRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type AdminIdRequest struct {
	Id int64 `uri:"id"`
}

type AdminCreateRequest struct {
	User_id int64 `json:"user_id"`
	Role_id int64 `json:"role_id"`
}

type AdminUpdateRequest struct {
	Id       int64  `json:"id"`
	User_id int64 `json:"user_id"`
	Role_id int64 `json:"role_id"`
}
