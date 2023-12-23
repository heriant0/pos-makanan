package auth

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthToken struct {
	Access_token string `json:"access_token"`
	Role         string `json:"role"`
}

type AuthTokenPayload struct {
	Id int `json:"id"`
	Email string `json:"email"`
}