package token

type ResultResponse struct {
	Name      string `json:"name"`
	Token     string `json:"token"`
	Role      string `json:"role"`
	ExpiredAt string `json:"expired_at"`
}
