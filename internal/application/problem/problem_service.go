package problem_app

import (
	problem_domain "plms_be/internal/domain/problem"
)

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

func (p *ProblemAppService) GetTestCaseById(id int64) (*problem_domain.TestCase, error) {
	TestCase, err := p.ProblemService.GetTestCaseByIdDomain(id)
	if err != nil {
		return nil, err
	}
	return TestCase, nil
}

func (p *ProblemAppService) SaveTestCase(testCase *problem_domain.PartialTestCaseUpdate)  error {
	err := p.ProblemService.SaveTestCaseDomain(testCase)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemAppService) SubmitProblem(submit *problem_domain.SubmitProblem) (*string , error) {
	job_id ,  err := p.ProblemService.SubmitProblemDomain(submit)
	if err != nil {
		return nil, err
	}
	return job_id, nil
}

func (p *ProblemAppService) CheckSubmissionStatus(job_id string) (*problem_domain.SubmissionResult, error) {
	result, err := p.ProblemService.CheckSubmissionStatusDomain(job_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}