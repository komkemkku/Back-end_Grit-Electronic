package requests

type AdminLogRequest struct {
	Page      int64  `form:"page"`
	Size      int64  `form:"size"`
	Search    string `form:"search"`
	AdminID   int    `form:"admin_id"`
	StartDate int64  `form:"start_date"`
	EndDate   int64  `form:"end_date"`
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
