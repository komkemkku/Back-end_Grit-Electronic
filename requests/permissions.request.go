package requests

type PermissionRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type PermissionIdRequest struct {
	Id int64 `uri:"id"`
}

type PermissionUpdateRequest struct {
	Id          int64  `json:"id"`
	Group_name  string `json:"group_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
