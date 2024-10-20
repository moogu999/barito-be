package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
)

func TestBeginTx(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				mockDB.ExpectBegin()
			},
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

			repo := NewOrderRepository(db)

			_, err = repo.BeginTx(ctx)

			if (err != nil) != tt.wantErr {
				t.Errorf("BeginTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommitTx(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				mockDB.ExpectBegin()
				mockDB.ExpectCommit()
			},
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

			repo := NewOrderRepository(db)

			tx, _ := repo.BeginTx(ctx)
			err = repo.CommitTx(tx)

			if (err != nil) != tt.wantErr {
				t.Errorf("CommitTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRollbackTx(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				mockDB.ExpectBegin()
				mockDB.ExpectRollback()
			},
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

			repo := NewOrderRepository(db)

			tx, _ := repo.BeginTx(ctx)
			err = repo.RollbackTx(tx)

			if (err != nil) != tt.wantErr {
				t.Errorf("RollbackTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	query := `INSERT INTO orders (user_id,total_amount,created_at) VALUES (?,?,?)`
	now := time.Now()
	order := &entity.Order{
		UserID:      1,
		TotalAmount: 50.0,
		CreatedAt:   now,
	}
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		order   *entity.Order
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectBegin()
				mockDB.ExpectExec(query).
					WithArgs(order.UserID, order.TotalAmount, order.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			order:   order,
			wantErr: false,
		},
		{
			name: "failed to execute",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectBegin()
				mockDB.ExpectExec(query).
					WillReturnError(err)
			},
			order:   order,
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

			repo := NewOrderRepository(db)

			tx, _ := repo.BeginTx(ctx)
			err = repo.CreateOrder(ctx, tx, tt.order)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOrderItem(t *testing.T) {
	t.Parallel()

	query := `INSERT INTO order_items (book_id,qty,price) VALUES (?,?,?)`
	item := &entity.OrderItem{
		BookID: 1,
		Qty:    1,
		Price:  10.0,
	}
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		item    *entity.OrderItem
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectBegin()
				mockDB.ExpectExec(query).
					WithArgs(item.BookID, item.Qty, item.Price).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			item:    item,
			wantErr: false,
		},
		{
			name: "failed to execute",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectBegin()
				mockDB.ExpectExec(query).
					WillReturnError(err)
			},
			item:    item,
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

			repo := NewOrderRepository(db)

			tx, _ := repo.BeginTx(ctx)
			err = repo.CreateOrderItem(ctx, tx, tt.item)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrderItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
