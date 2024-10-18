package repository

import (
	"context"

	"github.com/moogu999/barito-be/internal/domain/entity"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
}
