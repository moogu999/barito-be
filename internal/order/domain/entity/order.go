package entity

import (
	"time"
)

type Order struct {
	ID          int64
	UserID      int64
	Items       []OrderItem
	TotalAmount float64
	CreatedAt   time.Time
}

func NewOrder(userID int64, items []OrderItem) Order {
	itemsMap := make(map[int64]OrderItem)

	totalAmount := 0.0
	for _, val := range items {
		totalAmount += val.Price

		if item, ok := itemsMap[val.ID]; ok {
			item.Qty += val.Qty
			itemsMap[val.ID] = item
		} else {
			itemsMap[val.ID] = val
		}
	}

	groupedItems := make([]OrderItem, len(itemsMap))
	for _, val := range itemsMap {
		groupedItems = append(groupedItems, val)
	}

	return Order{
		UserID:      userID,
		Items:       groupedItems,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now().UTC(),
	}
}
