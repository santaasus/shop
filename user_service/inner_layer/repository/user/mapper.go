package user

import (
	domain "shop/user_service/inner_layer/domain/user"
	"time"
)

func MapUpdateUserToParams(user domain.UpdateUser) (pararms map[string]any) {
	pararms = map[string]any{
		"email":      user.Email,
		"user_name":  user.UserName,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"updated_at": time.Now().Format(time.DateTime),
	}

	for k, v := range pararms {
		if v == "" {
			delete(pararms, k)
		}
	}

	return
}

func MapToDomainUser(newUser *domain.NewUser) *domain.User {
	return &domain.User{
		Email:     newUser.Email,
		UserName:  newUser.UserName,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
	}
}
