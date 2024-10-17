package db

import (
	"database/sql"
	"fmt"
	"time"

	dbCore "shop/db_service"
	domain "shop/user_service/inner_layer/domain/user"
)

const (
	SQLSELECT = "SELECT * FROM users WHERE "
	SQLUPDATE = "UPDATE users SET "
)

func CreateUser(newUser *domain.User) (*domain.User, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var userId int
	query := `INSERT INTO users (email, hash_password, user_name, first_name, last_name, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	actualTime := time.Now()
	// https://github.com/golang/go/blob/master/src/time/format.go
	timeString := actualTime.Format(time.DateTime)
	err = db.QueryRow(
		query,
		newUser.Email, newUser.HashPassword,
		newUser.UserName, newUser.FirstName,
		newUser.LastName, timeString, timeString,
	).Scan(&userId)

	if err != nil {
		return nil, err
	}

	newUser.ID = userId
	newUser.CreatedAt = actualTime
	newUser.UpdatedAt = actualTime

	return newUser, nil
}

func GetUserByID(id int) (*domain.User, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT * FROM users WHERE id = $1;"

	row := db.QueryRow(query, id)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	var user domain.User

	err = scanToUser(row, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUserByID(id int) error {
	db, err := dbCore.Connect()
	if err != nil {
		return nil
	}

	defer db.Close()

	query := "DELETE FROM users WHERE id=$1;"
	_, err = db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByParams(params map[string]any) (*domain.User, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query, queryArgs := configureQueryBy(SQLSELECT, params)
	row := db.QueryRow(query, queryArgs...)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	var user domain.User

	if err := scanToUser(row, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserByParams(params map[string]any, userId int) (*domain.User, error) {
	db, err := dbCore.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query, queryArgs := configureQueryBy(SQLUPDATE, params)
	query += fmt.Sprintf(" WHERE id=%d", userId) + " RETURNING *;"

	row := db.QueryRow(query, queryArgs...)
	err = row.Err()
	if err != nil {
		return nil, err
	}

	var user domain.User

	if err := scanToUser(row, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func configureQueryBy(sqlAction string, params map[string]any) (string, []any) {
	query := sqlAction
	queryArgs := []any{}
	argumentCounter := 1

	// Assemble query, args by map arams. It's looks like:
	// SELECT * FROM users WHERE field1 = $1 AND field1 = $1
	for key, value := range params {
		if value == nil {
			continue
		}

		query += fmt.Sprintf("%s = $%d", key, argumentCounter)

		if len(params) != argumentCounter {
			if sqlAction == SQLUPDATE {
				query += ","
			} else {
				query += " AND "
			}
		}

		queryArgs = append(queryArgs, value)

		argumentCounter++
	}

	return query, queryArgs
}

func scanToUser(row *sql.Row, user *domain.User) error {
	if err := row.Scan(&user.ID, &user.UserName,
		&user.Email, &user.FirstName,
		&user.LastName, &user.HashPassword,
		&user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}

	return nil
}
