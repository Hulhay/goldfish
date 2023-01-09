package profile

type ProfileResponse struct {
	ID           int64  `json:"id"`
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
}
