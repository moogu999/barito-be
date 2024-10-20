package mock

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

type MockService struct {
	FindBooksFunc func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error)
}

func (m MockService) FindBooks(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
	if m.FindBooksFunc != nil {
		return m.FindBooksFunc(ctx, params)
	}

	return nil, nil
}
