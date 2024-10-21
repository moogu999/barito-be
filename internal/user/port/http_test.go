package port

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/user/domain/entity"
	"github.com/moogu999/barito-be/internal/user/port/oapi"
	"github.com/moogu999/barito-be/internal/user/usecase/mock"
	"github.com/oapi-codegen/runtime/types"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	request := oapi.NewUser{
		Email:    types.Email("testing@testing.com"),
		Password: "testing",
	}
	tests := []struct {
		name           string
		request        oapi.NewUser
		mockFunc       func(ctx context.Context, mockService *mock.MockService)
		wantStatusCode int
	}{
		{
			name:    "success",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateUserFunc = func(ctx context.Context, email, password string) error {
					return nil
				}
			},
			wantStatusCode: http.StatusCreated,
		},
		{
			name:    "error",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateUserFunc = func(ctx context.Context, email, password string) error {
					return errors.New("err")
				}
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:    "email is already being used",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateUserFunc = func(ctx context.Context, email, password string) error {
					return entity.ErrEmailIsUsed
				}
			},
			wantStatusCode: http.StatusConflict,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockService := mock.MockService{}
			tt.mockFunc(ctx, &mockService)

			r := chi.NewRouter()
			handler := NewHandler(r, mockService)

			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/users", bytes.NewBuffer(body))
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Fatalf("/v1/users = %v, wantStatusCode %v", rr.Code, tt.wantStatusCode)
			}
		})
	}
}

func TestCreateSession(t *testing.T) {
	t.Parallel()

	request := oapi.CreateUserRequestObject{
		Body: &oapi.NewUser{
			Email:    types.Email("testing@testing.com"),
			Password: "testing",
		},
	}
	tests := []struct {
		name           string
		request        oapi.CreateUserRequestObject
		mockFunc       func(ctx context.Context, mockService *mock.MockService)
		wantStatusCode int
	}{
		{
			name:    "success",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateSessionFunc = func(ctx context.Context, email, password string) (int64, error) {
					return 1, nil
				}
			},
			wantStatusCode: http.StatusCreated,
		},
		{
			name:    "error",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateSessionFunc = func(ctx context.Context, email, password string) (int64, error) {
					return 0, errors.New("err")
				}
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:    "email is not registered",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateSessionFunc = func(ctx context.Context, email, password string) (int64, error) {
					return 0, entity.ErrNotRegistered
				}
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:    "incorrect password",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.CreateSessionFunc = func(ctx context.Context, email, password string) (int64, error) {
					return 0, entity.ErrIncorrectPassword
				}
			},
			wantStatusCode: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockService := mock.MockService{}
			tt.mockFunc(ctx, &mockService)

			r := chi.NewRouter()
			handler := NewHandler(r, mockService)

			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/sessions", bytes.NewBuffer(body))
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Fatalf("/v1/sessions = %v, wantStatusCode %v", rr.Code, tt.wantStatusCode)
			}
		})
	}
}
