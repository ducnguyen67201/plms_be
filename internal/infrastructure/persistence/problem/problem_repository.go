package problem_db

import (
	"database/sql"
	"fmt"
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
			firstRow = false
		}

		problem.TestCase = append(problem.TestCase, problem_domain.TestCase{
			TestCaseID:     testCaseID,
			Input:          input,
			ExpectedOutput: expectedOutput,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			IsActive:       isActive,
		})
	}

	return problem, nil
}

func (r *OracleProblemRepository) GetProblemById(id int64) (*problem_domain.Problem, error) {
	query := `select * from problem WHERE problem_id = :1;`

	row := r.DB.QueryRow(query, int(id))

	var problem problem_domain.Problem
	err := row.Scan(
		&problem.ProblemID,
		&problem.ContestID,
		&problem.Title,
		&problem.Description,
		&problem.DifficultyLevel,
		&problem.RepeatedTimes,
		&problem.Type,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no problem found with ID %d", id)
		}
		return nil, err
	}

	return &problem, nil
}

func (r *OracleProblemRepository) SaveProblemDomain(problem *problem_domain.Problem) error {
	query := `UPDATE Problem
		SET contest_id = :1,
			title = :2,
			description = :3,
			difficulty_level = :4,
			repeated_times = :5,
			type = :6
		WHERE problem_id = :7;`

	_, err := r.DB.Exec(query,
		&problem.ContestID,
		problem.Title,
		problem.Description,
		problem.DifficultyLevel,
		problem.RepeatedTimes,
		problem.Type,
		problem.ProblemID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *OracleProblemRepository) GetTestCaseById(id int64) (*problem_domain.TestCase, error) {
	query := `SELECT * FROM Test_Case WHERE test_case_id = :1;`
	
	row := r.DB.QueryRow(query, id)
	var testCase problem_domain.TestCase
	err := row.Scan(
		&testCase.TestCaseID,
		&testCase.ProblemID,
		&testCase.Input,
		&testCase.ExpectedOutput,
		&testCase.CreatedAt,
		&testCase.UpdatedAt,
		&testCase.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no test case found with ID %d", id)
		}
		return nil, err
	}
	return &testCase, nil

}

func (r *OracleProblemRepository) SaveTestCaseDomain(testCase *problem_domain.TestCase) error {	
	query := `
	UPDATE Test_Case
	SET problem_id = :1,
		Input = :2,
		Expected_Output = :3,
		Created_At = TO_DATE(:4, 'YYYY-MM-DD HH24:MI:SS'),
		Updated_At = TO_DATE(:5, 'YYYY-MM-DD HH24:MI:SS'),
		Is_active = :6
	WHERE test_case_id = :7;`

	_, err := r.DB.Exec(query,
		testCase.ProblemID,
		testCase.Input,
		testCase.ExpectedOutput,
		testCase.CreatedAt,
		testCase.UpdatedAt,
		testCase.IsActive,
		testCase.TestCaseID,
	)

	if err != nil {
		return err
	}

	return nil
}