package service

import "github.com/romoguill/farming-backend/internal/model"

type UserRepository interface {
	GetMany() ([]model.User, error)
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.UserRepository.GetMany()
}
