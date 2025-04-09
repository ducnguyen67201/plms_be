package problem_db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	problem_domain "plms_be/internal/domain/problem"
	ViewModel "plms_be/viewModel"

	"github.com/go-redis/redis/v8"
)

type OracleProblemRepository struct {
	DB *sql.DB
	Redis *redis.Client
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
			&problem.MethodName,
			&problem.SkeletonCode,
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
	defer rows.Close()
	for rows.Next() { 
		var ( 
			problemID       int64
			contestID       *int64
			title           string
			description     string
			difficultyLevel string
			repeatedTimes   int64
			problemType     string
			methodName    string
			skeletonCode   string

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
			&methodName,
			&skeletonCode,

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
				MethodName:      methodName,
				SkeletonCode:    skeletonCode,
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

func (r *OracleProblemRepository) SubmitJobInProgress(submit *ViewModel.CodeJob) error {
	// * Write data to Redis,indicate job is in progress 
	err := r.Redis.Set(r.Redis.Context(), submit.JobID, `{"result": "in_progress"}`, 0).Err()
	if err != nil {
		log.Println("Error setting value in Redis:", err)
		return err
	}
	return nil
}

func (r *OracleProblemRepository) CheckSubmissionStatus(job_id string) (*problem_domain.SubmissionResult, error) {
	var result problem_domain.SubmissionResult
	// * Check for code submission status in Redis
	value , err := r.Redis.Get(r.Redis.Context(), job_id).Result()
	if err == redis.Nil { 
		return nil, fmt.Errorf("job id not found")
	} else if err != nil { 
		return nil , fmt.Errorf("error getting value from Redis: %v", err)
	} else { 
		log.Printf("value from Redis: %s", value)
	}

	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &result, nil
}