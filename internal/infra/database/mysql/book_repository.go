package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) repository.BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) FindBooks(ctx context.Context, params repository.BookFilter) ([]*entity.Book, error) {
	builder := sq.Select("id", "title", "author", "isbn", "price").
		From("books")

	if params.Author != "" {
		builder = builder.Where(sq.Eq{"author": params.Author})
	}

	if params.Title != "" {
		builder = builder.Where(sq.Eq{"title": params.Title})
	}

	q, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*entity.Book, 0)
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.Price,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func (r *BookRepository) GetBooksByIDs(ctx context.Context, ids []int64) ([]*entity.Book, error) {
	builder := sq.Select("id", "title", "author", "isbn", "price").
		From("books").
		Where(sq.Eq{"id": ids})

	q, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*entity.Book, 0)
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.ISBN,
			&book.Price,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}
