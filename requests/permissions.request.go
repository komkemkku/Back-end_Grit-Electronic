package requests

type PermissionRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type PermissionIdRequest struct {
	ID int `uri:"id"`
}

type PermissionUpdateRequest struct {
	Id          int    `json:"id"`
	GroupName   string `json:"group_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
