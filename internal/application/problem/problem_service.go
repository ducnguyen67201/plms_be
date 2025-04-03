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

func (p *ProblemAppService) GetProblemById(id string) (*problem_domain.ProblemWithTestCase, error) {
	problems, err := p.ProblemService.GetProblemByIdDomain(id)
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (p *ProblemAppService) SaveProblem(problem *problem_domain.PartialProblemUpdate) error {
	err := p.ProblemService.SaveProblemDomain(problem)
	if err != nil {
		return err
	}
	return nil
}