package discussion_db

import (
	"database/sql"
	"fmt"
	discussion_domain "plms_be/internal/domain/discussion"
	"time"
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

func (r *OracleDiscussionRepository) CreateCommentOnDiscussionPostId(input *discussion_domain.CreateDiscussionComent) error {
	stmt := `BEGIN CreateCommentOnDiscussionPostId(:discussion_id, :user_id, :user_comment); END;`

	var result string

	_, err := r.DB.Exec(stmt,
		input.DiscussionID,
		input.UserID,
		input.UserComment,
	)
	fmt.Printf("result: %s\n", result)
	if err != nil {
		return  err
	}

	return nil
}

func (r *OracleDiscussionRepository) GetAllCommentOnDiscussionPostId(id int64) (*discussion_domain.DiscussionWithComment, error) {
	query := `SELECT * FROM GetDiscussionComment WHERE discussion_id = :1;`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discussion *discussion_domain.DiscussionWithComment
	var firstRow bool = true

	for rows.Next() {
		var (
			discussionID    int64
			title           string
			topic           string
			content         string
			discussionLike  int64
			createdDate     time.Time
			createdBy       int64

			userID      int64
			userComment string
			username    string
		)

		err := rows.Scan(
			&discussionID,
			&title,
			&topic,
			&content,
			&discussionLike,
			&createdDate,
			&createdBy,
			&userID,
			&userComment,
			&username,
		)
		if err != nil {
			return nil, err
		}

		if firstRow {
			discussion = &discussion_domain.DiscussionWithComment{
				DiscussionID:    discussionID,
				Title:           title,
				Topic:           topic,
				Content:         content,
				Discussion_like: discussionLike,
				Created_date:    createdDate,
				Created_by:      createdBy,
				Comment:         []discussion_domain.Comment{},
			}
			firstRow = false
		}

		discussion.Comment = append(discussion.Comment, discussion_domain.Comment{
			UserID:      userID,
			UserComment: userComment,
			Username:    username,
		})
	}

	if discussion == nil {
		return nil, nil
	}

	return discussion, nil
}