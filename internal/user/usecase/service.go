package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/user/domain/repository"
)

type User interface {
	CreateUser(ctx context.Context, email, password string) error
}

type Service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) User {
	return &Service{
		repo: repo,
	}
}
