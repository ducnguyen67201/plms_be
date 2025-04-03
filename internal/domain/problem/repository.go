package problem_domain

type ProblemRepository interface {
	GetAllProblemDomain() ([]*Problem, error)
	GetProblemByIdDomain(id string) (*ProblemWithTestCase, error)
}
