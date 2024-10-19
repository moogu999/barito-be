package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
	"github.com/moogu999/barito-be/internal/book/domain/repository/mock"
)

func TestFindBooks(t *testing.T) {
	t.Parallel()

	filter := repository.BookFilter{
		Title:  "john",
		Author: "doe",
	}
	books := []*entity.Book{
		{
			ID:     1,
			Title:  "john",
			Author: "doe",
			ISBN:   "testing",
			Stock:  100,
		},
	}
	err := errors.New("err")

	tests := []struct {
		name     string
		filter   repository.BookFilter
		mockFunc func(ctx context.Context, mockRepo *mock.MockBookRepository)
		want     []*entity.Book
		wantErr  bool
	}{
		{
			name:   "success",
			filter: filter,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockBookRepository) {
				mockRepo.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return books, nil
				}
			},
			want:    books,
			wantErr: false,
		},
		{
			name:   "no results",
			filter: filter,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockBookRepository) {
				mockRepo.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return []*entity.Book{}, nil
				}
			},
			want:    []*entity.Book{},
			wantErr: false,
		},
		{
			name:   "repo.FindBooks error",
			filter: filter,
			mockFunc: func(ctx context.Context, mockRepo *mock.MockBookRepository) {
				mockRepo.FindBooksFunc = func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
					return nil, err
				}
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockRepo := mock.MockBookRepository{}
			tt.mockFunc(ctx, &mockRepo)

			service := NewService(mockRepo)

			got, err := service.FindBooks(ctx, tt.filter)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindBooks() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, got) && !tt.wantErr {
				t.Errorf("FindBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
