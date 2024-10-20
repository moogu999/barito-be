package repository

import (
	"context"

	"github.com/moogu999/barito-be/internal/user/domain/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
}
