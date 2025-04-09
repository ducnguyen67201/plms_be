package discussion_domain

type DiscussionRepository interface {
	GetAllDiscussion() ([]*Discussion, error)

	GetDiscussionById(id int64) (*Discussion, error)
	SaveDiscussion(discussion *Discussion) error

	CreateCommentOnDiscussionPostId(id *CreateDiscussionComent) error
	GetAllCommentOnDiscussionPostId(id int64) (*DiscussionWithComment, error)
}