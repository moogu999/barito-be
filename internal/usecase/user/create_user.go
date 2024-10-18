package user

import (
	"context"
	"errors"
	"strings"

	"github.com/moogu999/barito-be/internal/domain/entity"
)

func (s *Service) CreateUser(ctx context.Context, email, password string) error {
	existingUser, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existingUser != nil && strings.EqualFold(existingUser.Email, email) {
		return errors.New("email is already being used")
	}

	newUser, err := entity.NewUser(email, password)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(ctx, &newUser)
}
