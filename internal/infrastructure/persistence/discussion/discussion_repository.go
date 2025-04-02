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