package response

type PermissionResponses struct {
	Id          int  `json:"id"`
	Group_name  string `json:"group_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Create_at   int64  `json:"create_at"`
}
