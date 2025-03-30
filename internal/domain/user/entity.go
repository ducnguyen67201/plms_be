package user_domain

import "errors"

type User struct {
	UserID          int64  `json:"user_id"`
	RoleID          int64  `json:"role_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	RegistrationDate string `json:"registration_date"`
}

func NewUser(user_id int64, role_id int64, username string, password string, email string, registration_date string) (*User, error) {
	if user_id <= 0 || role_id <= 0 || username == "" || password == "" || email == "" || registration_date == "" {
		return nil, errors.New("invalid user data")
	}

	return &User{
		UserID:          user_id,
		RoleID:          role_id,
		Username:        username,
		Password:        password,
		Email:           email,
		RegistrationDate: registration_date,
	}, nil
}