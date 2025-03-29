package user_app

import (
	user_domain "plms_be/internal/domain/user"
)

type UserAppService struct {
	UserService *user_domain.Service
}

func (s *UserAppService) Register(user user_domain.User) (*user_domain.User, error) {
	domainUser, err := s.UserService.RegisterUser(1, 0, "", "", "", "")
	if err != nil {
		return nil, err
	}
	return domainUser, nil
}