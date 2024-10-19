package repository

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
)

type BookFilter struct {
	Title  string
	Author string
}

type BookRepostiroy interface {
	FindBooks(ctx context.Context, params BookFilter) ([]*entity.Book, error)
}
