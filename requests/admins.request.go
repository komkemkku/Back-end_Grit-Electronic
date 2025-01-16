package requests

type AdminRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type AdminIdRequest struct {
	ID int `uri:"id"`
}

type AdminCreateRequest struct {
	RoleID   int    `json:"role_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

type AdminUpdateRequest struct {
	ID       int    `json:"id"`
	RoleID   int    `json:"role_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}
