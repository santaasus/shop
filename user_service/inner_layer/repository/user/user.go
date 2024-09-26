package user

import (
	"encoding/json"
	"errors"
	rDB "shop/db_service/redis"
	db "shop/user_service/inner_layer/db"
	domainErrors "shop/user_service/inner_layer/domain/errors"
	domain "shop/user_service/inner_layer/domain/user"
	"strconv"

	_ "github.com/lib/pq"
)

type IRepository interface {
	GetUserByID(id int) (*domain.User, error)
	GetUserByParams(params map[string]any) (*domain.User, error)
	CreateUser(newUser *domain.User) (*domain.User, error)
	UpdateUser(updateUser domain.UpdateUser, userId int) error
	DeleteUserByID(userId int) error
}

type Repository struct {
}

func (Repository) GetUserByID(id int) (*domain.User, error) {
	value, _ := rDB.GetValueBy(rDB.USER + strconv.Itoa(id))
	if value != "" {
		var user domain.User
		err := json.Unmarshal([]byte(value), &user)

		return &user, err
	}

	user, err := db.GetUserByID(id)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("user does not exist"),
			Type: domainErrors.ValidationError,
		}
	}

	rDB.SaveBy(rDB.USER+strconv.Itoa(id), user)

	return user, nil
}

func (Repository) GetUserByParams(params map[string]any) (*domain.User, error) {
	user, err := db.GetUserByParams(params)
	if err != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("user does not exist"),
			Type: domainErrors.ValidationError,
		}
	}

	return user, nil
}

func (Repository) CreateUser(newUser *domain.User) (*domain.User, error) {
	existUser, _ := db.GetUserByParams(map[string]any{"email": newUser.Email})
	if existUser != nil {
		return nil, &domainErrors.AppError{
			Err:  errors.New("user already exist"),
			Type: domainErrors.ValidationError,
		}
	}
	user, err := db.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (Repository) UpdateUser(updateUser domain.UpdateUser, userId int) error {
	params := MapUpdateUserToParams(updateUser)

	if len(params) == 0 {
		return &domainErrors.AppError{
			Err:  errors.New("wrong the request body"),
			Type: domainErrors.ValidationError,
		}
	}

	err := db.UpdateUserByParams(params, userId)
	if err != nil {
		return err
	}

	return nil
}

func (Repository) DeleteUserByID(userId int) error {
	err := db.DeleteUserByID(userId)
	if err != nil {
		return err
	}

	rDB.DeleteValueBy(rDB.USER + strconv.Itoa(userId))

	return nil
}
