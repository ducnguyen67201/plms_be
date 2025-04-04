package learning_app

import learning_domain "plms_be/internal/domain/learning"

type LearningAppService struct {
	LearningService *learning_domain.LearningService
}

func (l *LearningAppService) GetAllLearning() ([]*learning_domain.LearningMaterial, error) {
	learningMaterials, err := l.LearningService.GetAllLearning()
	if err != nil {
		return nil, err
	}
	return learningMaterials, nil
}

func (l *LearningAppService) GetLearningByID(id int64) (*learning_domain.LearningMaterial, error) {
	learningMaterial, err := l.LearningService.GetLearningByID(id)
	if err != nil {
		return nil, err
	}
	return learningMaterial, nil
}

func (l *LearningAppService) SaveLearning(learningMaterial *learning_domain.PartialLearningMaterial) error {
	err := l.LearningService.SaveLearning(learningMaterial)
	if err != nil {
		return err
	}
	return nil
}