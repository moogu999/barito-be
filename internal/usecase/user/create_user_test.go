package user

import (
	"context"
	"errors"
	"testing"

	"github.com/moogu999/barito-be/internal/domain/entity"
	"github.com/moogu999/barito-be/internal/domain/mock"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	email := "testing@testing.com"
	password := "testing"
	err := errors.New("err")
	tests := []struct {
		name     string
		email    string
		password string
		mockFunc func(ctx context.Context, mockRepo *mock.MockUserRepository)
		wantErr  bool
	}{
		{
			name:     "success",
			email:    email,
			password: password,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return nil, nil
				}
				mockRepo.CreateUserFunc = func(ctx context.Context, user *entity.User) error { return nil }
			},
			wantErr: false,
		},
		{
			name:     "email is already being used",
			email:    email,
			password: password,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return &entity.User{Email: email}, nil
				}
			},
			wantErr: true,
		},
		{
			name:     "password is too long",
			email:    email,
			password: "7DA1LRHz7KsRCS0dvO5A1CvjE5jDXDh2Z9iPeN1741260y8a9K2ze738aJxOztz7TRQ8lBLdZ",
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return nil, nil
				}
			},
			wantErr: true,
		},
		{
			name:     "repo.GetUserByEmail error",
			email:    email,
			password: password,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return nil, err
				}
			},
			wantErr: true,
		},
		{
			name:     "repo.CreateUser error",
			email:    email,
			password: password,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return nil, nil
				}
				mockRepo.CreateUserFunc = func(ctx context.Context, user *entity.User) error { return err }
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockRepo := mock.MockUserRepository{}
			tt.mockFunc(ctx, &mockRepo)

			service := NewService(mockRepo)

			err := service.CreateUser(ctx, tt.email, tt.password)

			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateUser() expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("CreateUser() didn't expect error but got %v", err)
				}
			}
		})
	}
}
