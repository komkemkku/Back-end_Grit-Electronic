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
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

type AdminUpdateRequest struct {
	ID       int64  `json:"id"`
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}
