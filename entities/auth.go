package entities

type AuthRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type UserAuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
