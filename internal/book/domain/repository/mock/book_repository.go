package mock

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

type MockBookRepository struct {
	FindBooksFunc     func(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error)
	GetBooksByIDsFunc func(ctx context.Context, ids []int64) ([]*entity.Book, error)
}

func (m MockBookRepository) FindBooks(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
	if m.FindBooksFunc != nil {
		return m.FindBooksFunc(ctx, params)
	}

	return nil, nil
}

func (m MockBookRepository) GetBooksByIDs(ctx context.Context, ids []int64) ([]*entity.Book, error) {
	if m.GetBooksByIDsFunc != nil {
		return m.GetBooksByIDsFunc(ctx, ids)
	}

	return nil, nil
}
