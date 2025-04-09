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

type CreateDiscussionComent struct { 
	DiscussionID int64 `json:"discussion_id"`
	UserID int64 `json:"user_id"`
	UserComment  string `json:"user_comment"`
}

type DiscussionWithComment struct { 
	DiscussionID    int64     `json:"discussion_id"`
	Title           string    `json:"title"`
	Topic           string    `json:"topic"`
	Content         string    `json:"content"`
	Discussion_like int64     `json:"discussion_like"`
	Created_date    time.Time `json:"created_date"`
	Created_by      int64     `json:"created_by"`
	Comment []Comment `json:"comment"`
}

type Comment struct { 
	UserID int64 `json:"user_id"`
	UserComment string `json:"user_comment"`
	Username string `json:"username"`
}