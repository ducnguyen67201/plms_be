package user_oracle_db

import (
	"context"
	"database/sql"
	"errors"
	user_domain "plms_be/internal/domain/user"
)

type OracleUserRepository  struct {
	DB *sql.DB
}

func (r *OracleUserRepository) GetById(user_id int64) (*user_domain.User, error) {

	return nil, nil
}

func (r *OracleUserRepository) GetByUsername(username string) (*user_domain.User, error) {
	return nil, nil
}

func (r *OracleUserRepository) FindByUsernameAndPassword(username string, password string) (*user_domain.User, error) {
	query := "SELECT * FROM Users WHERE username = :1 AND password = :2"
	row := r.DB.QueryRowContext(context.Background(), query, username, password)

	var user user_domain.User
	err := row.Scan(&user.UserID, &user.RoleID, &user.Username, &user.Password, &user.Email, &user.RegistrationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err 
	}

	return &user, nil
}