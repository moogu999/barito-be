package repository

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
)

type BookFilter struct {
	Title  string
	Author string
}

type BookRepository interface {
	FindBooks(ctx context.Context, params BookFilter) ([]*entity.Book, error)
	GetBooksByIDs(ctx context.Context, ids []int64) ([]*entity.Book, error)
}
