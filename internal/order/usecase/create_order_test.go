package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	bookEntity "github.com/moogu999/barito-be/internal/book/domain/entity"
	mockBookRepo "github.com/moogu999/barito-be/internal/book/domain/repository/mock"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	"github.com/moogu999/barito-be/internal/order/domain/repository/mock"
	userEntity "github.com/moogu999/barito-be/internal/user/domain/entity"
	mockUserRepo "github.com/moogu999/barito-be/internal/user/domain/repository/mock"
)

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	var userID int64 = 1
	cartItems := []CartItem{
		{
			BookID: 1,
			Qty:    1,
		},
		{
			BookID: 1,
			Qty:    1,
		},
		{
			BookID: 2,
			Qty:    1,
		},
	}
	user := userEntity.User{
		Email: "testing@testing.com",
	}
	books := []*bookEntity.Book{
		{
			ID:    1,
			Price: 10.0,
		},
		{
			ID:    2,
			Price: 20.0,
		},
	}
	err := errors.New("err")

	tests := []struct {
		name      string
		userID    int64
		cartItems []CartItem
		mockFunc  func(ctx context.Context,
			mockOrderRepo *mock.MockOrderRepository,
			mockUserRepo *mockUserRepo.UserRepository,
			mockBookRepo *mockBookRepo.MockBookRepository)
		want    int64
		wantErr bool
	}{
		{
			name:      "success",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
				mockOrderRepo.BeginTxFunc = func(ctx context.Context) (*sql.Tx, error) {
					return &sql.Tx{}, nil
				}
				mockOrderRepo.CommitTxFunc = func(tx *sql.Tx) error {
					return nil
				}
				mockOrderRepo.CreateOrderFunc = func(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
					order.ID = 1
					return nil
				}
				mockOrderRepo.CreateOrderItemFunc = func(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error {
					return nil
				}
			},
			want:    1,
			wantErr: false,
		},
		{
			name:      "userRepo.GetUserByID error",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return nil, err
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "bookRepo.GetBooksByIDs error",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return nil, err
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "user is not found",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return nil, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "books are not found",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return []*bookEntity.Book{
						{
							ID:    1,
							Price: 10.0,
						},
					}, nil
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "orderRepo.BeginTx error",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
				mockOrderRepo.BeginTxFunc = func(ctx context.Context) (*sql.Tx, error) {
					return nil, err
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "orderRepo.CreateOrder error",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
				mockOrderRepo.BeginTxFunc = func(ctx context.Context) (*sql.Tx, error) {
					return &sql.Tx{}, nil
				}
				mockOrderRepo.RollbackTxFunc = func(tx *sql.Tx) error {
					return nil
				}
				mockOrderRepo.CreateOrderFunc = func(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
					return err
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "orderRepo.CreateOrderItem error",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
				mockOrderRepo.BeginTxFunc = func(ctx context.Context) (*sql.Tx, error) {
					return &sql.Tx{}, nil
				}
				mockOrderRepo.RollbackTxFunc = func(tx *sql.Tx) error {
					return nil
				}
				mockOrderRepo.CreateOrderFunc = func(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
					order.ID = 1
					return nil
				}
				mockOrderRepo.CreateOrderItemFunc = func(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error {
					return err
				}
			},
			want:    0,
			wantErr: true,
		},
		{
			name:      "orderRepo.CommitTx error",
			userID:    userID,
			cartItems: cartItems,
			mockFunc: func(ctx context.Context, mockOrderRepo *mock.MockOrderRepository, mockUserRepo *mockUserRepo.UserRepository, mockBookRepo *mockBookRepo.MockBookRepository) {
				mockUserRepo.GetUserByIDFunc = func(ctx context.Context, id int64) (*userEntity.User, error) {
					return &user, nil
				}
				mockBookRepo.GetBooksByIDsFunc = func(ctx context.Context, ids []int64) ([]*bookEntity.Book, error) {
					return books, nil
				}
				mockOrderRepo.BeginTxFunc = func(ctx context.Context) (*sql.Tx, error) {
					return &sql.Tx{}, nil
				}
				mockOrderRepo.CommitTxFunc = func(tx *sql.Tx) error {
					return err
				}
				mockOrderRepo.CreateOrderFunc = func(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
					order.ID = 1
					return nil
				}
				mockOrderRepo.CreateOrderItemFunc = func(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error {
					return nil
				}
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockOrderRepo := mock.MockOrderRepository{}
			mockUserRepo := mockUserRepo.UserRepository{}
			mockBookRepo := mockBookRepo.MockBookRepository{}
			tt.mockFunc(ctx, &mockOrderRepo, &mockUserRepo, &mockBookRepo)

			service := NewService(mockOrderRepo, mockUserRepo, mockBookRepo)

			got, err := service.CreateOrder(ctx, tt.userID, tt.cartItems)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.want != got {
				t.Errorf("CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
