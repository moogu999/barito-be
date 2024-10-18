package mysql

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moogu999/barito-be/internal/domain/entity"
)

func TestGetUserByEmail(t *testing.T) {
	t.Parallel()

	query := `SELECT id, email, password, created_at, created_by FROM users WHERE email = ?`
	email := "testing@testing.com"
	now := time.Now()
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		email   string
		want    *entity.User
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(email).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "created_at", "created_by"}).
						AddRow(1, email, "testing", now, email))
			},
			email: email,
			want: &entity.User{
				ID:        1,
				Email:     email,
				Password:  "testing",
				CreatedAt: now,
				CreatedBy: email,
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
			email:   email,
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to scan",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(email).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "created_at", "created_by"}).
						AddRow(1, email, "testing", nil, email))
			},
			email:   email,
			want:    nil,
			wantErr: true,
		},
		{
			name: "no match",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(email).
					WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "created_at", "created_by"}))
			},
			email:   email,
			want:    nil,
			wantErr: false,
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

			repo := NewUserRepository(db)

			got, err := repo.GetUserByEmail(ctx, tt.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, got) && !tt.wantErr {
				t.Errorf("GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	query := `INSERT INTO users (email,password,created_at,created_by) VALUES (?,?,?,?)`
	now := time.Now()
	user := &entity.User{
		Email:     "testing@testing.com",
		Password:  "testing",
		CreatedAt: now,
		CreatedBy: "testing@testing.com",
	}
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		user    *entity.User
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectExec(query).
					WithArgs(user.Email, user.Password, user.CreatedAt, user.CreatedBy).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			user:    user,
			wantErr: false,
		},
		{
			name: "failed to execute",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectExec(query).
					WillReturnError(err)
			},
			user:    user,
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

			repo := NewUserRepository(db)

			err = repo.CreateUser(ctx, tt.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
