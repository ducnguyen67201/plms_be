package user_domain

type Repository interface {
	GetById(user_id int64) (*User, error)
	GetByUsername(username string) (*User, error)
	FindByUsernameAndPassword(username string, password string) (*User, error)
}