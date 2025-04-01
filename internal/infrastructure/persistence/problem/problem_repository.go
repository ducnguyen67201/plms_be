package problem_db

import (
	"database/sql"
	problem_domain "plms_be/internal/domain/problem"
)

type OracleProblemRepository struct {
	DB *sql.DB
}

func (r *OracleProblemRepository) GetAllProblemDomain() ([]*problem_domain.Problem, error) {
	query := "SELECT * FROM Problem;"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []*problem_domain.Problem
	for rows.Next() {
		var problem problem_domain.Problem
		if err := rows.Scan(
			&problem.ProblemID,
			&problem.ContestID,
			&problem.Title,
			&problem.Description,
			&problem.DifficultyLevel,
			&problem.RepeatedTimes,
			&problem.Type,
		); err != nil {
			return nil, err
		}
		problems = append(problems, &problem)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return problems, nil
}