package problem_domain

import "time"

type Problem struct {
	ProblemID       int64  `json:"problem_id"`
	ContestID       *int64 `json:"contest_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	DifficultyLevel string `json:"difficulty_level"`
	RepeatedTimes   int64  `json:"repeated_times"`
	Type            string `json:"type"`
	MethodName    	string `json:"method_name"`
	SkeletonCode 	string `json:"skeleton_code"`
}

type PartialProblemUpdate struct {
	ProblemID       *int64  `json:"problem_id"`
	ContestID       *int64  `json:"contest_id"`
	Title           *string `json:"title"`
	Description     *string `json:"description"`
	DifficultyLevel *string `json:"difficulty_level"`
	RepeatedTimes   *int64  `json:"repeated_times"`
	Type            *string `json:"type"`
}

type ProblemWithTestCase struct {
	ProblemID       int64      `json:"problem_id"`
	ContestID       *int64     `json:"contest_id"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	DifficultyLevel string     `json:"difficulty_level"`
	RepeatedTimes   int64      `json:"repeated_times"`
	Type            string     `json:"type"`
	MethodName     string     `json:"method_name"`
	SkeletonCode  string     `json:"skeleton_code"`

	TestCase        []TestCase `json:"test_cases"`
}

type TestCase struct {
	TestCaseID     int64   `json:"test_case_id"`
	ProblemID      int64   `json:"problem_id"`
	Input          string  `json:"input"`
	ExpectedOutput string  `json:"expected_output"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
	IsActive       string  `json:"is_active"`
}

type PartialTestCaseUpdate struct {
	TestCaseID     *int64  `json:"test_case_id"`
	ProblemID      *int64  `json:"problem_id"`
	Input          *string `json:"input"`
	ExpectedOutput *string `json:"expected_output"`
	CreatedAt      *string `json:"created_at"`
	UpdatedAt      *string `json:"updated_at"`
	IsActive       *string `json:"is_active"`
}

type SubmitProblem struct {
	SubmissionID   string    `json:"submission_id"`
	UserID 	  int64    `json:"user_id"`
	ProblemID 	int64     `json:"problem_id"`
	SubmissionDate time.Time `json:"submission_date"`
	Result         string    `json:"result"`
	Performance    string    `json:"performance"`
	Code 		 string    `json:"code"`
	Language       string    `json:"language"`
}

type SubmissionResult  struct { 
	JobID string `json:"job_id"`
	SubmissionID string `json:"submission_id"`
	UserID int64 `json:"user_id"`
	ProblemID int64 `json:"problem_id"`
	SubmissionDate time.Time `json:"submission_date"`
	Result string `json:"result"`
	Performance string `json:"performance"`
}