package entity

import (
	"time"
)

// Order represents an order in the system
type Order struct {
	ID          int64
	UserID      int64
	Email       string
	Items       []OrderItem
	TotalAmount float64
	CreatedAt   time.Time
}

// New order return a new order.
// Itgroups together items with the same bookID
// and calculates the total amount for the order.
// If the item quantity is less than 1, it will return an error.
func NewOrder(userID int64, items []OrderItem) (Order, error) {
	itemsMap := make(map[int64]OrderItem)

	totalAmount := 0.0
	for _, val := range items {
		if val.Qty < 1 {
			return Order{}, ErrInvalidQuantity
		}

		totalAmount += val.Price

		if item, ok := itemsMap[val.BookID]; ok {
			item.Qty += val.Qty
			itemsMap[val.BookID] = item
		} else {
			itemsMap[val.BookID] = val
		}
	}

	groupedItems := make([]OrderItem, 0)
	for _, val := range itemsMap {
		groupedItems = append(groupedItems, val)
	}

	return Order{
		UserID:      userID,
		Items:       groupedItems,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now().UTC(),
	}, nil
}
