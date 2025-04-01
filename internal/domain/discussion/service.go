package discussion_domain

type DiscussionService struct {
	repo DiscussionRepository
}

func NewDiscussionService(discussionRepository DiscussionRepository) *DiscussionService {
	return &DiscussionService{repo: discussionRepository}
}

func (r *DiscussionService) GetAllDiscussion() ([]*Discussion, error) {
	discussions, err := r.repo.GetAllDiscussion()
	if err != nil {
		return nil, err
	}

	return discussions, nil
}