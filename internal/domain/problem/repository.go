package problem_domain

type ProblemRepository interface {
	GetAllProblemDomain() ([]*Problem, error)
}