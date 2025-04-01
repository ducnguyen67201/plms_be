package user_domain

import "errors"

type User struct {
	UserID          int64  `json:"user_id"`
	RoleID          int64  `json:"role_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	RegistrationDate string `json:"registration_date"`
	Profile Profile `json:"profile"`
	Role string `json:"role"`
}

type Profile struct { 
	Profile_id int64 `json:"profile_id"`
	User_id int64 `json:"user_id"`
	Bio string `json:"bio"`
	Picture_url string `json:"picture_url"`
	Location string `json:"location"`
	Linkedin_url string `json:"linkedin_url"`
	Github_url string `json:"github_url"`
	Profile_view int64 `json:"profile_view"`
	Contest_attending int64 `json:"contest_attending"`
	Problem_solved int64 `json:"problem_solved"`
	Date_of_birth string `json:"date_of_birth"`
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