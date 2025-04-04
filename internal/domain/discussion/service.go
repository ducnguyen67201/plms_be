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

func (r *DiscussionService) GetDiscussionById(id int64) (*Discussion, error) {
	discussion, err := r.repo.GetDiscussionById(id)
	if err != nil {
		return nil, err
	}

	return discussion, nil
}

func (r *DiscussionService) SaveDiscussion(discussion *PartialDiscussionUpdate) error {
	discussionByID, err := r.repo.GetDiscussionById(*discussion.DiscussionID)
	if err != nil {
		return err
	}

	if discussion.Title != nil {
		discussionByID.Title = *discussion.Title
	}
	if discussion.Topic != nil {
		discussionByID.Topic = *discussion.Topic
	}
	if discussion.Content != nil {
		discussionByID.Content = *discussion.Content
	}
	if discussion.Discussion_like != nil {
		discussionByID.Discussion_like = *discussion.Discussion_like
	}
	if discussion.Created_date != nil {
		discussionByID.Created_date = *discussion.Created_date
	}

	if discussion.Created_by != nil {
		discussionByID.Created_by = *discussion.Created_by
	}

	err = r.repo.SaveDiscussion(discussionByID)
	if err != nil {
		return err
	}

	return nil
}