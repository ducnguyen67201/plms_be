package learning_domain

type LearningRepository interface {
	GetAllLearning() ([]*LearningMaterial, error)
}