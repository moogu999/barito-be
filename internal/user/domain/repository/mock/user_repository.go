package mock

import (
	"context"

	"github.com/moogu999/barito-be/internal/user/domain/entity"
)

// @TODO rename
type UserRepository struct {
	GetUserByEmailFunc func(ctx context.Context, email string) (*entity.User, error)
	CreateUserFunc     func(ctx context.Context, user *entity.User) error
	GetUserByIDFunc    func(ctx context.Context, id int64) (*entity.User, error)
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

func (m UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(ctx, id)
	}

	return nil, nil
}
