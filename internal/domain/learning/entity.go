package learning_domain

type LearningMaterial struct {
	MaterialID         string `json:"material_id"`
	MaterialCategoryID string `json:"material_category_id"`
	Posted_by          string `json:"posted_by"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	MateriaDate        string `json:"material_date"`
}