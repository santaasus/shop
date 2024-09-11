package auth

type LoginRequest struct {
	Email    string `json:"email" example:"test@gmail.com" binding:"required"`
	Password string `json:"password" example:"Test1234" binding:"required"`
}

type AccessTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
