package requests

type ReportRequest struct {
	Page   int64  `form:"page"`
	Size   int64  `form:"size"`
	Search string `form:"search"`
}

type ReportIdRequest struct {
	ID int64 `uri:"id"`
}