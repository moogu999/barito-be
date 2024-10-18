package user

import "github.com/moogu999/barito-be/internal/domain/repository"

type Service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}
