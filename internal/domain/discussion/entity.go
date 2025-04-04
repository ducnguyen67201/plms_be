package discussion_domain

import "time"

type Discussion struct {
	DiscussionID    int64     `json:"discussion_id"`
	Title           string    `json:"title"`
	Topic           string    `json:"topic"`
	Content         string    `json:"content"`
	Discussion_like int64     `json:"discussion_like"`
	Created_date    time.Time `json:"created_date"`
	Created_by      int64     `json:"created_by"`
}

type PartialDiscussionUpdate struct {
	DiscussionID    *int64  `json:"discussion_id"`
	Title           *string `json:"title"`
	Topic           *string `json:"topic"`
	Content         *string `json:"content"`
	Discussion_like *int64  `json:"discussion_like"`
	Created_date    *time.Time `json:"created_date"`
	Created_by      *int64  `json:"created_by"`
}