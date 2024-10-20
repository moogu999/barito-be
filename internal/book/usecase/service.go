package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

type BookUseCase interface {
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
