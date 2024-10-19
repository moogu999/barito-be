package mock

import "context"

type Service struct {
	CreateUserFunc func(ctx context.Context, email, password string) error
}

func (m Service) CreateUser(ctx context.Context, email, password string) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(ctx, email, password)
	}

	return nil
}
