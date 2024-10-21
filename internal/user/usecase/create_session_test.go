package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/moogu999/barito-be/internal/user/domain/entity"
	"github.com/moogu999/barito-be/internal/user/domain/repository/mock"
)

func TestCreateSession(t *testing.T) {
	t.Parallel()

	email := "testing@testing.com"
	password := "testing"
	user, _ := entity.NewUser(email, password)
	user.ID = 1
	err := errors.New("err")

	tests := []struct {
		name     string
		email    string
		password string
		mockFunc func(ctx context.Context, mockRepo *mock.MockUserRepository)
		want     int64
		wantErr  bool
	}{
		{
			name:     "success",
			email:    email,
			password: password,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return &user, nil
				}
			},
			want:    1,
			wantErr: false,
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
			want:    0,
			wantErr: true,
		},
		{
			name:     "email is not registered",
			email:    email,
			password: password,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return nil, nil
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:     "incorrect password",
			email:    email,
			password: "1234567890",
			mockFunc: func(ctx context.Context, mockRepo *mock.MockUserRepository) {
				mockRepo.GetUserByEmailFunc = func(ctx context.Context, email string) (*entity.User, error) {
					return &user, nil
				}
			},
			want:    0,
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

			got, err := service.CreateSession(ctx, tt.email, tt.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSession() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("CreateSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
