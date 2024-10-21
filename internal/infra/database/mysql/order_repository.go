package mysql

import (
	"context"
	"database/sql"
	"time"

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

func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID int64) ([]*entity.Order, error) {
	builder := sq.
		Select(
			"o.id",
			"o.user_id",
			"u.email",
			"i.id AS item_id",
			"i.book_id",
			"b.title",
			"b.author",
			"i.qty",
			"i.price",
			"o.total_amount",
			"o.created_at").
		From("orders o").
		Join("users u ON o.user_id = u.id").
		Join("order_items i ON o.id = i.order_id").
		Join("books b ON i.book_id = b.id").
		Where(sq.Eq{"o.user_id": userID})

	q, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int64]*entity.Order)
	for rows.Next() {
		var (
			id          int64
			userID      int64
			email       string
			itemID      int64
			bookID      int64
			title       string
			author      string
			qty         int
			price       float64
			totalAmount float64
			createdAt   time.Time
		)

		err = rows.Scan(
			&id,
			&userID,
			&email,
			&itemID,
			&bookID,
			&title,
			&author,
			&qty,
			&price,
			&totalAmount,
			&createdAt,
		)
		if err != nil {
			return nil, err
		}

		if _, ok := ordersMap[id]; !ok {
			ordersMap[id] = &entity.Order{
				ID:          id,
				UserID:      userID,
				Email:       email,
				Items:       []entity.OrderItem{},
				TotalAmount: totalAmount,
				CreatedAt:   createdAt,
			}
		}

		ordersMap[id].Items = append(ordersMap[id].Items, entity.OrderItem{
			ID:     itemID,
			BookID: bookID,
			Title:  title,
			Author: author,
			Qty:    qty,
			Price:  price,
		})
	}

	orders := make([]*entity.Order, 0)
	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	return orders, nil
}
