package usecase

import (
	"context"
	"strings"

	"github.com/moogu999/barito-be/internal/user/domain/entity"
)

func (s *Service) CreateUser(ctx context.Context, email, password string) error {
	existingUser, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existingUser != nil && strings.EqualFold(existingUser.Email, email) {
		return entity.ErrEmailIsUsed
	}

	newUser, err := entity.NewUser(email, password)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(ctx, &newUser)
}
