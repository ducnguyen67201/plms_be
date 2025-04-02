package learning_domain

type LearningService struct {
	repo LearningRepository
}

func NewLearningService(learningRepository LearningRepository) *LearningService {
	return &LearningService{repo: learningRepository}
}

func (r *LearningService) GetAllLearning() ([]*LearningMaterial, error) {
	learningMaterials, err := r.repo.GetAllLearning()
	if err != nil {
		return nil, err
	}

	return learningMaterials, nil
}
