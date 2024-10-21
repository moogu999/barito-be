package port

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	bookEntity "github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	"github.com/moogu999/barito-be/internal/order/port/oapi"
	"github.com/moogu999/barito-be/internal/order/usecase"
	"github.com/moogu999/barito-be/internal/order/usecase/mock"
	userEntity "github.com/moogu999/barito-be/internal/user/domain/entity"
)

func TestCreateOrder(t *testing.T) {
	t.Parallel()

	request := oapi.NewOrder{
		UserId: 1,
		Items: []oapi.Item{
			{
				BookId: 1,
				Qty:    10,
			},
		},
	}
	tests := []struct {
		name           string
		request        oapi.NewOrder
		mockFunc       func(ctx context.Context, mock *mock.MockService)
		wantStatusCode int
	}{
		{
			name:    "success",
			request: request,
			mockFunc: func(ctx context.Context, mock *mock.MockService) {
				mock.CreateOrderFunc = func(ctx context.Context, userID int64, items []usecase.CartItem) (int64, error) {
					return 1, nil
				}
			},
			wantStatusCode: http.StatusCreated,
		},
		{
			name:    "error",
			request: request,
			mockFunc: func(ctx context.Context, mock *mock.MockService) {
				mock.CreateOrderFunc = func(ctx context.Context, userID int64, items []usecase.CartItem) (int64, error) {
					return 0, errors.New("err")
				}
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:    "user is not found",
			request: request,
			mockFunc: func(ctx context.Context, mock *mock.MockService) {
				mock.CreateOrderFunc = func(ctx context.Context, userID int64, items []usecase.CartItem) (int64, error) {
					return 0, userEntity.ErrUserNotFound
				}
			},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:    "user is not found",
			request: request,
			mockFunc: func(ctx context.Context, mock *mock.MockService) {
				mock.CreateOrderFunc = func(ctx context.Context, userID int64, items []usecase.CartItem) (int64, error) {
					return 0, bookEntity.ErrBooksNotFound
				}
			},
			wantStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mock := mock.MockService{}
			tt.mockFunc(ctx, &mock)

			r := chi.NewRouter()
			handler := NewHandler(r, mock)

			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/orders", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Fatalf("/v1/orders = %v, wantStatusCode %v", rr.Code, tt.wantStatusCode)
			}
		})
	}
}

func TestFindOrders(t *testing.T) {
	t.Parallel()

	var userID int64 = 1
	tests := []struct {
		name           string
		userID         int64
		mockFunc       func(ctx context.Context, mock *mock.MockService)
		wantStatusCode int
	}{
		{
			name:   "success",
			userID: userID,
			mockFunc: func(ctx context.Context, mock *mock.MockService) {
				mock.FindOrdersFunc = func(ctx context.Context, userID int64) ([]*entity.Order, error) {
					return []*entity.Order{
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
						},
					}, nil
				}
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name:   "error",
			userID: userID,
			mockFunc: func(ctx context.Context, mock *mock.MockService) {
				mock.FindOrdersFunc = func(ctx context.Context, userID int64) ([]*entity.Order, error) {
					return nil, errors.New("err")
				}
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mock := mock.MockService{}
			tt.mockFunc(ctx, &mock)

			r := chi.NewRouter()
			handler := NewHandler(r, mock)

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/orders?userId=%d", tt.userID), nil)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Fatalf("/v1/orders = %v, wantStatusCode %v", rr.Code, tt.wantStatusCode)
			}
		})
	}
}
