package port

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
	"github.com/moogu999/barito-be/internal/book/port/oapi"
	"github.com/moogu999/barito-be/internal/book/usecase/mock"
)

func TestFindBooks(t *testing.T) {
	t.Parallel()

	title := "john"
	author := "doe"
	request := oapi.FindBooksRequestObject{
		Params: oapi.FindBooksParams{
			Title:  &title,
			Author: &author,
		},
	}
	tests := []struct {
		name           string
		request        oapi.FindBooksRequestObject
		mockFunc       func(ctx context.Context, mockService *mock.MockService)
		wantStatusCode int
	}{
		{
			name:    "success",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return []*entity.Book{
						{
							ID:     1,
							Title:  "John",
							Author: "Doe",
							ISBN:   "testing",
							Price:  100.0,
						},
					}, nil
				}
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:    "error",
			request: request,
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return nil, errors.New("err")
				}
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "empty author",
			request: oapi.FindBooksRequestObject{
				Params: oapi.FindBooksParams{
					Title: &title,
				},
			},
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return nil, nil
				}
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "empty title",
			request: oapi.FindBooksRequestObject{
				Params: oapi.FindBooksParams{
					Author: &author,
				},
			},
			mockFunc: func(ctx context.Context, mockService *mock.MockService) {
				mockService.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return nil, nil
				}
			},
			wantStatusCode: http.StatusOK,
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

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/books?author=%s&title=%s", author, title), nil)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Fatalf("/v1/books = %v, wantStatusCode %v", rr.Code, tt.wantStatusCode)
			}
		})
	}
}
