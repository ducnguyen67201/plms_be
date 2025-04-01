package problem_app

import problem_domain "plms_be/internal/domain/problem"

type ProblemAppService struct {
	ProblemService *problem_domain.ProblemService
}

func (p *ProblemAppService) GetAllProblem() ([]*problem_domain.Problem, error) {
	problems, err := p.ProblemService.GetAllProblemDomain()
	if err != nil {
		return nil, err
	}
	return problems, nil
}