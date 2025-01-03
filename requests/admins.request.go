package requests

type AdminRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type AdminIdRequest struct {
	ID int64 `uri:"id"`
}

type AdminCreateRequest struct {
	Name     string  `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleID   int64  `json:"role_id"`
}

type AdminUpdateRequest struct {
	ID       int64  `json:"id"`
	Name     string  `json:"usnameer_id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleID   int64  `json:"role_id"`
}
