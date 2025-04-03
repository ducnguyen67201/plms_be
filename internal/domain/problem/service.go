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

func (s *ProblemService) SaveProblemDomain(patch *PartialProblemUpdate) error {
	problemByID, err := s.repo.GetProblemById(*patch.ProblemID)
	if err != nil {
		return err
	}

	if patch.Title != nil {
		problemByID.Title = *patch.Title
	}
	if patch.Description != nil {
		problemByID.Description = *patch.Description
	}
	if patch.DifficultyLevel != nil {
		problemByID.DifficultyLevel = *patch.DifficultyLevel
	}
	if patch.RepeatedTimes != nil {
		problemByID.RepeatedTimes = *patch.RepeatedTimes
	}
	if patch.Type != nil {
		problemByID.Type = *patch.Type
	}

	err = s.repo.SaveProblemDomain(problemByID)
	if err != nil {
		return err
	}
	return nil
}
