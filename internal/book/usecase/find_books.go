package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

func (s *Service) FindBooks(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
	books, err := s.repo.FindBooks(ctx, params)
	if err != nil {
		return nil, err
	}

	return books, nil
}
