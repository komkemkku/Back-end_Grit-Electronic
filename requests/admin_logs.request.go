package requests

type AdminLogRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type AdminLogIdRequest struct {
	ID int `uri:"id"`
}

type AdminLogCreateRequest struct {
	AdminID     int    `json:"admin_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
}

type AdminLogUpdateRequest struct {
	ID          int    `json:"id"`
	AdminID     int    `json:"admin_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
}
