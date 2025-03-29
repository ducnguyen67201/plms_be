package user_oracle_db

import (
	"database/sql"
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
