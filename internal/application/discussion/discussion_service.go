package discussion_app

import discussion_domain "plms_be/internal/domain/discussion"

type DiscussionAppService struct {
	DiscussionService *discussion_domain.DiscussionService
}

func (d *DiscussionAppService) GetAllDiscussion() ([]*discussion_domain.Discussion, error) {
	discussions, err := d.DiscussionService.GetAllDiscussion()
	if err != nil {
		return nil, err
	}
	return discussions, nil
}

func (d *DiscussionAppService) GetDiscussionById(id int64) (*discussion_domain.Discussion, error) {
	discussion, err := d.DiscussionService.GetDiscussionById(id)
	if err != nil {
		return nil, err
	}
	return discussion, nil
}

func (d *DiscussionAppService) SaveDiscussion(discussion *discussion_domain.PartialDiscussionUpdate)  error {
	err := d.DiscussionService.SaveDiscussion(discussion)
	if err != nil {
		return err
	}
	return nil
}

func (d *DiscussionAppService) CreateCommentOnDiscussionPostId(input *discussion_domain.CreateDiscussionComent) error {
	err := d.DiscussionService.CreateCommentOnDiscussionPostId(input)
	if err != nil {
		return err
	}
	return nil
}

func (d *DiscussionAppService) GetAllCommentOnDiscussionPostId(id int64) (*discussion_domain.DiscussionWithComment, error) {
	comments, err := d.DiscussionService.GetAllCommentOnDiscussionPostId(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}