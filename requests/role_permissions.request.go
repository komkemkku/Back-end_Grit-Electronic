package requests

type RolePermissionRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type RolePermissionIdRequest struct {
	ID int `uri:"id"`
}

type RolePremissionUpdateRequest struct {
	ID           int  `json:"id"`
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}
