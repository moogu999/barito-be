package mock

import (
	"context"

	"github.com/moogu999/barito-be/internal/domain/entity"
)

type UserRepository struct {
	GetUserByEmailFunc func(ctx context.Context, email string) (*entity.User, error)
	CreateUserFunc     func(ctx context.Context, user *entity.User) error
}

func (m UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(ctx, email)
	}

	return nil, nil
}

func (m UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(ctx, user)
	}

	return nil
}
