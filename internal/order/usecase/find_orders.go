package usecase

import (
	"context"

	"github.com/moogu999/barito-be/internal/order/domain/entity"
)

func (s *Service) FindOrders(ctx context.Context, userID int64) ([]*entity.Order, error) {
	orders, err := s.orderRepo.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
