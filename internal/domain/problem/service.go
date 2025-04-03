package problem_domain

type ProblemService struct {
	repo ProblemRepository
}

func NewProblemService(repo ProblemRepository) *ProblemService {
	return &ProblemService{repo: repo}
}

func (s *ProblemService) GetAllProblemDomain() ([]*Problem, error) {
	problems, err := s.repo.GetAllProblemDomain()
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (s *ProblemService) GetProblemByIdDomain(id string) (*ProblemWithTestCase, error) {
	problem, err := s.repo.GetProblemByIdDomain(id)
	if err != nil {
		return nil, err
	}
	return problem, nil
}
