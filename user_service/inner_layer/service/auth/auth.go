package auth

import (
	"errors"
	errorsDomain "github.com/santaasus/errors-handler"
	domain "shop/user_service/inner_layer/domain/user"
	repository "shop/user_service/inner_layer/repository/user"
	security "shop/user_service/inner_layer/security"
	"time"
)

type Service struct {
	UserRepository repository.IRepository
}

func (s *Service) Login(login *domain.LoginUser) (*AuthenticatedUser, error) {
	if login.Email == "" {
		err := &errorsDomain.AppError{
			Err:  errors.New("email is empty"),
			Type: errorsDomain.NotFound,
		}
		return nil, err
	}

	userParams := map[string]any{"email": login.Email}
	domainUser, err := s.UserRepository.GetUserByParams(userParams)
	if err != nil {
		return nil, err
	}

	if domainUser.ID == 0 {
		return &AuthenticatedUser{},
			&errorsDomain.AppError{
				Err:  errors.New("email or password does not match"),
				Type: errorsDomain.Unauthorized,
			}
	}

	isAuthenticated := security.IsFineCheckPasswordAndHash(domainUser.HashPassword, login.Password)
	if !isAuthenticated {
		return &AuthenticatedUser{},
			&errorsDomain.AppError{
				Err:  errors.New("email or password does not match"),
				Type: errorsDomain.Unauthorized,
			}
	}

	accessToken, err := security.GenerateJWTToken(domainUser.ID, security.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := security.GenerateJWTToken(domainUser.ID, security.Refresh)
	if err != nil {
		return nil, err
	}

	return secAuthUserMapper(domainUser, &SecurityData{
		JWTAccessToken:            accessToken.Token,
		JWTRefreshToken:           refreshToken.Token,
		ExpirationAccessDateTime:  accessToken.ExpirationTime,
		ExpirationRefreshDateTime: refreshToken.ExpirationTime,
	}), nil
}

func (s *Service) AccessTokenByRefreshToken(refreshToken string) (*AuthenticatedUser, error) {
	claims, err := security.VerifyTokenAndGetClaims(refreshToken, "refresh")
	if err != nil {
		return nil, err
	}

	domainUser, err := s.UserRepository.GetUserByID(claims["id"].(int))
	if err != nil {
		return &AuthenticatedUser{},
			&errorsDomain.AppError{
				Err:  errors.New("user `id` does not match"),
				Type: errorsDomain.NotFound,
			}
	}

	accessToken, err := security.GenerateJWTToken(domainUser.ID, "access")
	if err != nil {
		return nil, err
	}

	nextExpTime := int64(claims["exp"].(float64))

	return secAuthUserMapper(domainUser, &SecurityData{
		JWTAccessToken:            accessToken.Token,
		JWTRefreshToken:           refreshToken,
		ExpirationAccessDateTime:  accessToken.ExpirationTime,
		ExpirationRefreshDateTime: time.Unix(nextExpTime, 0),
	}), nil
}
