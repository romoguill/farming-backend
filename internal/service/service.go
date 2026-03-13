package service

type Repository interface {
	UserRepository
}

type Service struct {
	repository  Repository
	userService *UserService
}

func NewService(repository Repository) *Service {
	return &Service{
		repository:  repository,
		userService: NewUserService(repository),
	}
}

func (s *Service) UserService() *UserService {
	return s.userService
}
