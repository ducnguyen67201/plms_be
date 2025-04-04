package discussion_db

import (
	"database/sql"
	discussion_domain "plms_be/internal/domain/discussion"
)

type OracleDiscussionRepository struct {
	DB *sql.DB
}

func (r *OracleDiscussionRepository) GetAllDiscussion() ([]*discussion_domain.Discussion, error) {
	rows, err := r.DB.Query("SELECT * FROM Discussion;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discussions []*discussion_domain.Discussion
	for rows.Next() {
		var discussion discussion_domain.Discussion
		if err := rows.Scan(&discussion.DiscussionID, &discussion.Title, &discussion.Topic, &discussion.Content, &discussion.Discussion_like, &discussion.Created_date, &discussion.Created_by); err != nil {
			return nil, err
		}
		discussions = append(discussions, &discussion)
	}
	return discussions, nil
}

func (r *OracleDiscussionRepository) GetDiscussionById(id int64) (*discussion_domain.Discussion, error) {
	row := r.DB.QueryRow("SELECT * FROM Discussion WHERE discussion_id = :1", id)

	var discussion discussion_domain.Discussion
	if err := row.Scan(&discussion.DiscussionID, &discussion.Title, &discussion.Topic, &discussion.Content, &discussion.Discussion_like, &discussion.Created_date, &discussion.Created_by); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return &discussion, nil
}

func (r *OracleDiscussionRepository) SaveDiscussion(discussion *discussion_domain.Discussion) error {
	query := `
	UPDATE Discussion
	SET 
		title = :1, 
		topic = :2, 
		content = :3, 
		discussion_like = :4, 
		created_date = :5,
		created_by = :6
	WHERE discussion_id = :7`

	_, err := r.DB.Exec(query, 
		discussion.Title, 
		discussion.Topic,
		discussion.Content, 
		discussion.Discussion_like, 
		discussion.Created_date, 
		discussion.Created_by,
		discussion.DiscussionID,
	)
	if err != nil {
		return err
	}
	return nil
}