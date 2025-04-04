package learning_db

import (
	"database/sql"
	learning_domain "plms_be/internal/domain/learning"
	"time"
)

type OracleLearningRepository struct {
	DB *sql.DB
}

func (r *OracleLearningRepository) GetAllLearning() ([]*learning_domain.LearningMaterial, error) {
	rows, err := r.DB.Query("SELECT * FROM Learning_Material;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var learningMaterials []*learning_domain.LearningMaterial
	for rows.Next() {
		var learningMaterial learning_domain.LearningMaterial
		if err := rows.Scan(
			&learningMaterial.MaterialID, 
			&learningMaterial.MaterialCategoryID, 
			&learningMaterial.Posted_by, 
			&learningMaterial.Title, 
			&learningMaterial.Content, 
			&learningMaterial.MateriaDate); err != nil {
			return nil, err
		}

		if learningMaterial.MateriaDate != "" {
			parsedDate, err := time.Parse("2006-01-02T15:04:05-07:00", learningMaterial.MateriaDate)
			if err != nil {
				return nil, err
			}
			formattedDate := parsedDate.Format("2006-01-02 15:04:05")
			learningMaterial.MateriaDate = formattedDate
		}
		learningMaterials = append(learningMaterials, &learningMaterial)
		}
	return learningMaterials, nil
}

func (r *OracleLearningRepository) GetLearningByID(id int64) (*learning_domain.LearningMaterial, error) {
	query := `SELECT * FROM Learning_Material WHERE Material_ID = :1;`
	row := r.DB.QueryRow(query, id)

	var learningMaterial learning_domain.LearningMaterial
	if err := row.Scan(
		&learningMaterial.MaterialID, 
		&learningMaterial.MaterialCategoryID, 
		&learningMaterial.Posted_by, 
		&learningMaterial.Title, 
		&learningMaterial.Content, 
		&learningMaterial.MateriaDate); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}

	if learningMaterial.MateriaDate != "" {
		parsedDate, err := time.Parse("2006-01-02T15:04:05-07:00", learningMaterial.MateriaDate)
		if err != nil {
			return nil, err
		}
		formattedDate := parsedDate.Format("2006-01-02 15:04:05")
		learningMaterial.MateriaDate = formattedDate
	}
	return &learningMaterial, nil
}

func (r *OracleLearningRepository) SaveLearning(learningMaterial *learning_domain.LearningMaterial) error {
	query :=`
	UPDATE Learning_Material
	SET 
		Material_Category_ID = :1,
		Posted_by = :2,
		Title = :3,
		Content = :4,
		Material_Date = TO_DATE(:5, 'YYYY-MM-DD HH24:MI:SS')
	WHERE Material_ID = :6;`
	_, err := r.DB.Exec(query, 
		learningMaterial.MaterialCategoryID, 
		learningMaterial.Posted_by, 
		learningMaterial.Title, 
		learningMaterial.Content, 
		learningMaterial.MateriaDate,
		learningMaterial.MaterialID) 
	if err != nil {
		return err
	}
	return nil
}