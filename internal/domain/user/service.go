package user_domain

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterUser(userId int64, roleId int64, username string, password string, email string, registrationDate string) (*User, error) {
	user, err := NewUser(userId, roleId, username, password, email, registrationDate)
	if err != nil {
		return nil, err
	}

	// err = s.repo.Save(user)
	// if err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func (s *Service) Login(username string, password string) (*User, error) {
	user, err := s.repo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}