package requests

type RoleRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type RoleIdRequest struct {
	Id int64 `uri:"id"`
}

type RoleUpdateRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}