package user_domain

import "errors"

type User struct {
	user_id           int64  `json:"user_id"`
	role_id           int64  `json:"role_id"`
	username          string `json:"username"`
	password          string `json:"password"`
	email             string `json:"email"`
	registration_date string `json:"registration_date"`
}

func NewUser(user_id int64, role_id int64, username string, password string, email string, registration_date string) (*User, error) {
	if user_id <= 0 || role_id <= 0 || username == "" || password == "" || email == "" || registration_date == "" {
		return nil, errors.New("invalid user data")
	}

	return &User{
		user_id:           user_id,
		role_id:           role_id,
		username:          username,
		password:          password,
		email:             email,
		registration_date: registration_date,
	}, nil
}