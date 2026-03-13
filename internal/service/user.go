package service

import "github.com/romoguill/farming-backend/internal/model"

type UserRepository interface {
	GetUsers() ([]model.User, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repository.GetUsers()
}
