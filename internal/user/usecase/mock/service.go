package mock

import "context"

type MockService struct {
	CreateUserFunc    func(ctx context.Context, email, password string) error
	CreateSessionFunc func(ctx context.Context, email, password string) (int64, error)
}

func (m MockService) CreateUser(ctx context.Context, email, password string) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(ctx, email, password)
	}

	return nil
}

func (m MockService) CreateSession(ctx context.Context, email, password string) (int64, error) {
	if m.CreateSessionFunc != nil {
		return m.CreateSessionFunc(ctx, email, password)
	}

	return 0, nil
}
