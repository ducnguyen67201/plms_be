package problem_domain

import ViewModel "plms_be/viewModel"

type ProblemRepository interface {
	GetAllProblemDomain() ([]*Problem, error)
	GetProblemByIdDomain(id string) (*ProblemWithTestCase, error)
	SaveProblemDomain(problem *Problem) error
	GetProblemById(id int64) (*Problem, error)

	GetTestCaseById(id int64) (*TestCase, error)
	SaveTestCaseDomain(testCase *TestCase) error

	SubmitJobInProgress(job *ViewModel.CodeJob) error
	CheckSubmissionStatus(job_id string) (*SubmissionResult, error)
}
