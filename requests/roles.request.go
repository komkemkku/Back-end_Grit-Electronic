package requests

type RoleRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type RoleIdRequest struct {
	ID int `uri:"id"`
}

type RoleUpdateRequest struct {
	ID   int  `json:"id"`
	Name string `json:"name"`
}