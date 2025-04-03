package problem_db

import (
	"database/sql"
	"log"
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

func (r *OracleProblemRepository) GetProblemByIdDomain(id string) (*problem_domain.ProblemWithTestCase, error) {
	rows, err := r.DB.Query(`SELECT * FROM ProblemWithTestCases WHERE problem_id = :1`, id)

	if err != nil {
		log.Fatal("error executing query:", err)
		return nil, err
	}
	var problem *problem_domain.ProblemWithTestCase
	var firstRow bool = true 
	for rows.Next() { 
		var ( 
			problemID       int64
			contestID       *int64
			title           string
			description     string
			difficultyLevel string
			repeatedTimes   int64
			problemType     string

			testCaseID 	int64
			input 		string
			expectedOutput string
			createdAt 	*string
			updatedAt 	*string
			isActive 	string
		)

		err := rows.Scan(
			&problemID,
			&contestID,
			&title,
			&description,
			&difficultyLevel,
			&repeatedTimes,
			&problemType,

			&testCaseID,
			&input,
			&expectedOutput,
			&createdAt,
			&updatedAt,
			&isActive,
		)
		if err != nil {
			log.Fatal("error scanning row:", err)
		}

		if firstRow { 
			problem = &problem_domain.ProblemWithTestCase{
				ProblemID:       problemID,
				ContestID:       contestID,
				Title:           title,
				Description:     description,
				DifficultyLevel: difficultyLevel,
				RepeatedTimes:   repeatedTimes,
				Type:            problemType,
				TestCase:       []problem_domain.TestCase{},
			}
		}

		if isActive == "Y" { 
			problem.TestCase = append(problem.TestCase, problem_domain.TestCase{
				TestCaseID:     testCaseID,
				Input:          input,
				ExpectedOutput: expectedOutput,
				CreatedAt:      createdAt,
				UpdatedAt:      updatedAt,
				IsActive:       isActive,
			})
		}
	}

	return problem, nil
}

