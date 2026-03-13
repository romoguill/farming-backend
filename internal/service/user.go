package service

import "github.com/romoguill/farming-backend/internal/model"

type UserRepository interface {
	GetUsers() ([]model.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.userRepository.GetUsers()
}
