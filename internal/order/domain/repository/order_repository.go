package repository

import (
	"context"
	"database/sql"

	"github.com/moogu999/barito-be/internal/order/domain/entity"
)

type OrderRepository interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error

	CreateOrder(ctx context.Context, tx *sql.Tx, order *entity.Order) error

	CreateOrderItem(ctx context.Context, tx *sql.Tx, item *entity.OrderItem) error
}
