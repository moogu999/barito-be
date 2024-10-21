package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	mockBookRepo "github.com/moogu999/barito-be/internal/book/domain/repository/mock"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	"github.com/moogu999/barito-be/internal/order/domain/repository/mock"
	mockUserRepo "github.com/moogu999/barito-be/internal/user/domain/repository/mock"
)

func TestFindOrders(t *testing.T) {
	t.Parallel()

	var userID int64 = 1
	orders := []*entity.Order{
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
			CreatedAt:   time.Now(),
		},
	}
	err := errors.New("err")

	tests := []struct {
		name     string
		userID   int64
		mockFunc func(ctx context.Context, mock *mock.MockOrderRepository)
		want     []*entity.Order
		wantErr  bool
	}{
		{
			name:   "success",
			userID: userID,
			mockFunc: func(ctx context.Context, mock *mock.MockOrderRepository) {
				mock.GetOrdersByUserIDFunc = func(ctx context.Context, userID int64) ([]*entity.Order, error) {
					return orders, nil
				}
			},
			want:    orders,
			wantErr: false,
		},
		{
			name:   "orderRepo.GetOrdersByUserID error",
			userID: userID,
			mockFunc: func(ctx context.Context, mock *mock.MockOrderRepository) {
				mock.GetOrdersByUserIDFunc = func(ctx context.Context, userID int64) ([]*entity.Order, error) {
					return nil, err
				}
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockOrderRepo := mock.MockOrderRepository{}
			mockUserRepo := mockUserRepo.MockUserRepository{}
			mockBookRepo := mockBookRepo.MockBookRepository{}
			tt.mockFunc(ctx, &mockOrderRepo)

			service := NewService(mockOrderRepo, mockUserRepo, mockBookRepo)

			got, err := service.FindOrders(ctx, tt.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("FindOrders() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
