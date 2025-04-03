package problem_domain

type ProblemRepository interface {
	GetAllProblemDomain() ([]*Problem, error)
	GetProblemByIdDomain(id string) (*ProblemWithTestCase, error)
	SaveProblemDomain(problem *Problem) error
	GetProblemById(id int64) (*Problem, error)
}
