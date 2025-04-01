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
	query := `
		SELECT 
			u.user_id, u.role_id, u.username, u.password, u.email, u.registration_date, ur.role_name, 
			p.profile_id, p.user_id, p.bio, p.picture_url, p.linkedin_url, p.github_url, p.profile_view, 
			p.contest_attending, p.problem_solved, p.date_of_birth
		FROM Users u 
			JOIN User_Role ur ON u.role_id = ur.role_id
			JOIN Profile p ON u.user_id = p.user_id
		WHERE username = :1 AND password = :2;`

	row := r.DB.QueryRowContext(context.Background(), query, username, password)

	var user user_domain.User
	var profile user_domain.Profile
	
	err := row.Scan(
		&user.UserID,             
		&user.RoleID,             
		&user.Username,           
		&user.Password,           
		&user.Email,              
		&user.RegistrationDate, 
		&user.Role,                
		&profile.Profile_id,      
		&profile.User_id,         
		&profile.Bio,             
		&profile.Picture_url,     
		&profile.Linkedin_url,    
		&profile.Github_url,      
		&profile.Profile_view,    
		&profile.Contest_attending,
		&profile.Problem_solved, 
		&profile.Date_of_birth,  
	  )
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	user.Profile = profile

		return &user, nil
	}