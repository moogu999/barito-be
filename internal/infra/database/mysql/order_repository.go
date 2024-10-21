package mysql

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	"github.com/moogu999/barito-be/internal/order/domain/repository"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *OrderRepository) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (r *OrderRepository) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

func (r *OrderRepository) CreateOrder(ctx context.Context, tx *sql.Tx, order *entity.Order) error {
	builder := sq.Insert("orders").
		Columns("user_id", "total_amount", "created_at").
		Values(order.UserID, order.TotalAmount, order.CreatedAt)
	q, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	order.ID = id

	return nil
}

func (r *OrderRepository) CreateOrderItem(ctx context.Context, tx *sql.Tx, orderID int64, item *entity.OrderItem) error {
	builder := sq.Insert("order_items").
		Columns("order_id", "book_id", "qty", "price").
		Values(orderID, item.BookID, item.Qty, item.Price)
	q, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	item.ID = id

	return nil
}
