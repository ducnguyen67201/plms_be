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