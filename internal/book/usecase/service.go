package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

type BookUseCase interface {
	// FindBooks return the books available in the system.
	// They can be filtered by author and title.
	FindBooks(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error)
}

type Service struct {
	repo repository.BookRepository
}

func NewService(repo repository.BookRepository) *Service {
	return &Service{
		repo: repo,
	}
}
