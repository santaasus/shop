package auth

import "time"

// LoginUser is a struct that contains the request body for the login user
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	UserName  string `json:"userName" example:"UserName" gorm:"unique"`
	Email     string `json:"email" example:"some@mail.com" gorm:"unique"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
	Status    bool   `json:"status" example:"1"`
	Role      string `json:"role" example:"admin"`
	ID        int    `json:"id" example:"123"`
}

type SecurityData struct {
	JWTRefreshToken           string    `json:"jwtRefreshToken"`
	JWTAccessToken            string    `json:"jwtAccessToken"`
	ExpirationAccessDateTime  time.Time `json:"expirationAccessDateTime" example:"2023-02-02T21:03:53.196419-06:00"`
	ExpirationRefreshDateTime time.Time `json:"expirationRefreshDateTime" example:"2023-02-03T06:53:53.196419-06:00"`
}

type AuthenticatedUser struct {
	Data     UserData
	Security SecurityData
}
