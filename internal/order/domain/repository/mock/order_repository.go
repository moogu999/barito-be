package mock

import (
	"context"
	"database/sql"

	"github.com/moogu999/barito-be/internal/order/domain/entity"
)

type MockOrderRepository struct {
	BeginTxFunc         func(ctx context.Context) (*sql.Tx, error)
	CommitTxFunc        func(tx *sql.Tx) error
	RollbackTxFunc      func(tx *sql.Tx) error
	CreateOrderFunc     func(ctx context.Context, tx *sql.Tx, order *entity.Order) error
	CreateOrderItemFunc func(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error
}

func (m MockOrderRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	if m.BeginTxFunc != nil {
		return m.BeginTxFunc(ctx)
	}

	return nil, nil
}

func (m MockOrderRepository) CommitTx(tx *sql.Tx) error {
	if m.CommitTxFunc != nil {
		return m.CommitTxFunc(tx)
	}

	return nil
}

func (m MockOrderRepository) RollbackTx(tx *sql.Tx) error {
	if m.RollbackTxFunc != nil {
		return m.RollbackTxFunc(tx)
	}

	return nil
}

func (m MockOrderRepository) CreateOrder(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
	if m.CreateOrderFunc != nil {
		return m.CreateOrderFunc(ctx, tx, order)
	}

	return nil
}

func (m MockOrderRepository) CreateOrderItem(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error {
	if m.CreateOrderItemFunc != nil {
		return m.CreateOrderItemFunc(ctx, tx, orderID, item)
	}

	return nil
}
