package discussion_domain

type DiscussionRepository interface {
	GetAllDiscussion() ([]*Discussion, error)
}