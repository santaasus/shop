package auth

import (
	domain "shop/user_service/inner_layer/domain/user"
)

func secAuthUserMapper(domain *domain.User, oAuthInfo *SecurityData) *AuthenticatedUser {
	return &AuthenticatedUser{
		Data: UserData{
			ID:        domain.ID,
			UserName:  domain.UserName,
			Email:     domain.Email,
			FirstName: domain.FirstName,
			LastName:  domain.LastName,
			Status:    true,
		},
		Security: *oAuthInfo,
	}
}
