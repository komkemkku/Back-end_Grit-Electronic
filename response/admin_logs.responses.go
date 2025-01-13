package response

type AdminLogResponses struct {
	ID          int          `json:"id"`
	Admin       AdminLogResp `json:"admin"`
	Action      string       `json:"action"`
	Description string       `json:"description"`
	Created_at  int64        `json:"created_at"`
}
