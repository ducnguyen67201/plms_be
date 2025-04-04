package learning_domain

import "time"

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

func (r *LearningService) GetLearningByID(id int64) (*LearningMaterial, error) {
	learningMaterial, err := r.repo.GetLearningByID(id)
	if err != nil {
		return nil, err
	}

	return learningMaterial, nil
}

func (r *LearningService) SaveLearning(learningMaterial *PartialLearningMaterial) error {
	learningMaterialById, err := r.repo.GetLearningByID(*learningMaterial.MaterialID)
	if err != nil {
		return err
	}

	if learningMaterial.MaterialCategoryID != nil {
		learningMaterialById.MaterialCategoryID = *learningMaterial.MaterialCategoryID
	}
	if learningMaterial.Posted_by != nil {
		learningMaterialById.Posted_by = *learningMaterial.Posted_by
	}
	if learningMaterial.Title != nil {
		learningMaterialById.Title = *learningMaterial.Title
	}
	if learningMaterial.Content != nil {
		learningMaterialById.Content = *learningMaterial.Content
	}
	if learningMaterial.MateriaDate != nil {
		if learningMaterial.MateriaDate != nil {
			parsedDate, err := time.Parse(time.RFC3339, learningMaterial.MateriaDate.Format(time.RFC3339))
			if err != nil {
				return err
			}
			learningMaterialById.MateriaDate = parsedDate.Format(time.RFC3339)
		}
	}

	err = r.repo.SaveLearning(learningMaterialById)
	if err != nil {
		return err
	}

	return nil
}