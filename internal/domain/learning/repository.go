package learning_domain

type LearningRepository interface {
	GetAllLearning() ([]*LearningMaterial, error)

	GetLearningByID(id int64) (*LearningMaterial, error)
	SaveLearning(learningMaterial *LearningMaterial) error
}