package mysql

import (
	"context"
	"errors"
	"reflect"
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

	query := `INSERT INTO order_items (order_id,book_id,qty,price) VALUES (?,?,?,?)`
	var orderID int64 = 1
	item := &entity.OrderItem{
		BookID: 1,
		Qty:    1,
		Price:  10.0,
	}
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		orderID int64
		item    *entity.OrderItem
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectBegin()
				mockDB.ExpectExec(query).
					WithArgs(orderID, item.BookID, item.Qty, item.Price).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			orderID: orderID,
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
			orderID: orderID,
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
			err = repo.CreateOrderItem(ctx, tx, tt.orderID, tt.item)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrderItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOrdersByUserID(t *testing.T) {
	t.Parallel()

	query := `SELECT o.id, o.user_id, u.email, i.id AS item_id, i.book_id, b.title, b.author, i.qty, i.price, o.total_amount, o.created_at FROM orders o JOIN users u ON o.user_id = u.id JOIN order_items i ON o.id = i.order_id JOIN books b ON i.book_id = b.id WHERE o.user_id = ?`
	var userID int64 = 1
	now := time.Now()
	err := errors.New("err")

	tests := []struct {
		name    string
		setup   func(mockDB sqlmock.Sqlmock)
		userID  int64
		want    []*entity.Order
		wantErr bool
	}{
		{
			name: "success",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(userID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "email", "item_id", "book_id", "title", "author", "qty", "price", "total_amount", "created_at"}).
						AddRow(1, 1, "testing@testing.com", 1, 1, "John", "Doe", 1, 10.0, 10.0, now))
			},
			userID: userID,
			want: []*entity.Order{
				{
					ID:     1,
					UserID: 1,
					Email:  "testing@testing.com",
					Items: []entity.OrderItem{
						{
							ID:     1,
							BookID: 1,
							Title:  "John",
							Author: "Doe",
							Qty:    1,
							Price:  10.0,
						},
					},
					TotalAmount: 10.0,
					CreatedAt:   now,
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
			userID:  userID,
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to scan",
			setup: func(mockDB sqlmock.Sqlmock) {
				query := query

				mockDB.ExpectQuery(query).
					WithArgs(userID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "email", "item_id", "book_id", "title", "author", "qty", "price", "total_amount", "created_at"}).
						AddRow(1, 1, "testing@testing.com", 1, 1, "John", "Doe", 1, 10.0, nil, now))
			},
			userID:  userID,
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

			repo := NewOrderRepository(db)

			got, err := repo.GetOrdersByUserID(ctx, tt.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrdersByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, got) && !tt.wantErr {
				t.Errorf("GetOrdersByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
