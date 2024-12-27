package response

type AdminResponses struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
}
