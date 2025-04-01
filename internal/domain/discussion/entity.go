package discussion_domain

type Discussion struct {
	DiscussionID    string `json:"discussion_id"`
	Title           string `json:"title"`
	Topic           string `json:"topic"`
	Content         string `json:"content"`
	Discussion_like int64  `json:"discussion_like"`
	Created_date    string `json:"created_date"`
	Created_by      string `json:"created_by"`
}