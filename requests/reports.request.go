package requests

type ReportRequest struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
	// Search string `form:"search"`
	Month string `form:"month"`
	Year  int    `form:"year"`
}

type ReportIdRequest struct {
	ID int `uri:"id"`
}
