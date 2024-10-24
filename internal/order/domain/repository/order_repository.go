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
	GetOrdersByUserID(ctx context.Context, userID int64) ([]*entity.Order, error)

	CreateOrderItem(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error
}
