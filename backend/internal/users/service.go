package users

import "errors"

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
