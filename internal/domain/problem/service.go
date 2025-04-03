package problem_domain

import "time"

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

func (s *ProblemService) GetTestCaseByIdDomain(id int64) (*TestCase, error) {
	testCase, err := s.repo.GetTestCaseById(id)
	if err != nil {
		return nil, err
	}
	return testCase, nil
}

func (s *ProblemService) SaveTestCaseDomain(testCase *PartialTestCaseUpdate) error {
	testCaseById, err := s.repo.GetTestCaseById(*testCase.TestCaseID)
	if err != nil {
		return err
	}

	if testCase.ProblemID != nil {
		testCaseById.ProblemID = *testCase.ProblemID
	}
	if testCase.Input != nil {
		testCaseById.Input = *testCase.Input
	}
	if testCase.ExpectedOutput != nil {
		testCaseById.ExpectedOutput = *testCase.ExpectedOutput
	}
	if testCase.CreatedAt != nil {
		parsed, _ := time.Parse(time.RFC3339, *testCase.CreatedAt)
		formatted := parsed.Format("2006-01-02 15:04:05")
		testCaseById.CreatedAt = &formatted
	}

	updatedAt := time.Now().Format("2006-01-02 15:04:05")
	testCaseById.UpdatedAt = &updatedAt

	if testCase.IsActive != nil {
		testCaseById.IsActive = *testCase.IsActive
	}

	err = s.repo.SaveTestCaseDomain(testCaseById)
	if err != nil {
		return err
	}

	return nil
}