package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/user/domain/entity"
)

func (s *Service) CreateSession(ctx context.Context, email, password string) (int64, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, entity.ErrNotRegistered
	}

	if !user.VerifyPassword(password) {
		return 0, entity.ErrIncorrectPassword
	}

	return user.ID, nil
}
