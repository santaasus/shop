// Package domain contains the business logic for user/auth entities
package user

import "time"

type LoginUser struct {
	Email    string
	Password string
}

type NewUser struct {
	Email     string
	Password  string
	UserName  string
	FirstName string
	LastName  string
}

type UpdateUser struct {
	Email     string
	UserName  string
	FirstName string
	LastName  string
}

type User struct {
	ID           int
	UserName     string
	Email        string
	FirstName    string
	LastName     string
	HashPassword string
	CreatedAt    time.Time `example:"2024-07-26 05:23:20"`
	UpdatedAt    time.Time
}
