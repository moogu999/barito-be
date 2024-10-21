package mock

import (
	"context"

	"github.com/moogu999/barito-be/internal/order/usecase"
)

type MockService struct {
	CreateOrderFunc func(ctx context.Context, userID int64, items []usecase.CartItem) (int64, error)
}

func (m MockService) CreateOrder(ctx context.Context, userID int64, items []usecase.CartItem) (int64, error) {
	if m.CreateOrderFunc != nil {
		return m.CreateOrderFunc(ctx, userID, items)
	}

	return 0, nil
}
