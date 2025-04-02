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