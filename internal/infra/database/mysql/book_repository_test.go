package mysql

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/book/domain/repository"
)

func TestFindBooks(t *testing.T) {
	t.Parallel()

	query := `SELECT id, title, author, isbn, price FROM books WHERE author = ? AND title = ?`
	title := "john"
	author := "doe"
	filter := repository.BookFilter{
		Title:  title,
		Author: author,
	}
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		filter  repository.BookFilter
		want    []*entity.Book
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(author, title).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "isbn", "price"}).
						AddRow(1, title, author, "testing", 100.0))
			},
			filter: filter,
			want: []*entity.Book{
				{
					ID:     1,
					Title:  title,
					Author: author,
					ISBN:   "testing",
					Price:  100.0,
				},
			},
			wantErr: false,
		},
		{
			name: "failed to query",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WillReturnError(err)
			},
			filter:  filter,
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to scan",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(author, title).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "isbn", "price"}).
						AddRow(1, title, author, "testing", nil))
			},
			filter:  filter,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatal("error mocking sql")
			}
			defer db.Close()

			tt.setup(mock)

			repo := NewBookRepository(db)

			got, err := repo.FindBooks(ctx, tt.filter)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindBooks() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, got) && !tt.wantErr {
				t.Errorf("FindBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBooksByIDs(t *testing.T) {
	t.Parallel()

	query := `SELECT id, title, author, isbn, price FROM books WHERE id IN (?,?)`
	title := "john"
	author := "doe"
	ids := []int64{1, 2}
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		ids     []int64
		want    []*entity.Book
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "isbn", "price"}).
						AddRow(1, title, author, "testing", 10.0).
						AddRow(2, title, author, "testing", 30.0))
			},
			ids: ids,
			want: []*entity.Book{
				{
					ID:     1,
					Title:  title,
					Author: author,
					ISBN:   "testing",
					Price:  10.0,
				},
				{
					ID:     2,
					Title:  title,
					Author: author,
					ISBN:   "testing",
					Price:  30.0,
				},
			},
			wantErr: false,
		},
		{
			name: "failed to query",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WillReturnError(err)
			},
			ids:     ids,
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to scan",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(1, 2).
					WillReturnRows(sqlmock.NewRows([]string{"id", "title", "author", "isbn", "price"}).
						AddRow(1, title, author, "testing", nil).
						AddRow(2, title, author, "testing", 30.0))
			},
			ids:     ids,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatal("error mocking sql")
			}
			defer db.Close()

			tt.setup(mock)

			repo := NewBookRepository(db)

			got, err := repo.GetBooksByIDs(ctx, ids)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooksByIDs() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, got) && !tt.wantErr {
				t.Errorf("GetBooksByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
