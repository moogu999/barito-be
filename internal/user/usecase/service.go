package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/user/domain/repository"
)

type UserUseCase interface {
	// CreateUser registered a new user.
	// If an email is already registered to another user, it will return an error.
	CreateUser(ctx context.Context, email, password string) error

	// CreateSession logged in a user. It will return the user ID.
	// If the user is not found, it will return an error.
	CreateSession(Ctx context.Context, email, password string) (int64, error)
}

type Service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) UserUseCase {
	return &Service{
		repo: repo,
	}
}
