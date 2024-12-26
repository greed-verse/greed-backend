package jwt

type JwtPayload struct {
	Sub    string `json:"sub"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func EncodeJwt(claims ...string) string {
	return ""
}
