package learning_domain

import "time"

type LearningMaterial struct {
	MaterialID         int64    `json:"material_id"`
	MaterialCategoryID string `json:"material_category_id"`
	Posted_by          int64 `json:"posted_by"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	MateriaDate        string `json:"material_date"`
}

type PartialLearningMaterial struct {
	MaterialID         *int64    `json:"material_id"`
	MaterialCategoryID *string `json:"material_category_id"`
	Posted_by          *int64 `json:"posted_by"`
	Title              *string `json:"title"`
	Content            *string `json:"content"`
	MateriaDate        *time.Time `json:"material_date"`
}