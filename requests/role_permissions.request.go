package requests

type RolePermissionRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type RolePermissionIdRequest struct {
	Id int64 `uri:"id"`
}

type RolePremissionUpdateRequest struct {
	Id           int64  `json:"id"`
	Role_id       string `json:"role_id"`
	Permission_id string `json:"permission_id"`
}
