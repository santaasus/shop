package user

type NewUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUserRequest struct {
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserRequest struct {
	ID int `json:"id"`
}
