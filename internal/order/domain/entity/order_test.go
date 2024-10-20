package entity

import (
	"fmt"
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		userID int64
		items  []OrderItem
		want   Order
	}{
		{
			name:   "success",
			userID: 1,
			items: []OrderItem{
				{
					BookID: 1,
					Qty:    1,
					Price:  10.0,
				},
				{
					BookID: 1,
					Qty:    1,
					Price:  10.0,
				},
				{
					BookID: 2,
					Qty:    1,
					Price:  30.0,
				},
			},
			want: Order{
				UserID: 1,
				Items: []OrderItem{
					{
						BookID: 1,
						Qty:    2,
						Price:  10.0,
					},
					{
						BookID: 2,
						Qty:    1,
						Price:  30.0,
					},
				},
				TotalAmount: 50.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NewOrder(tt.userID, tt.items)

			if len(got.Items) != len(tt.want.Items) {
				t.Errorf("len(NewOrder().Items) = %v, want %v", len(got.Items), len(tt.want.Items))
			}

			if fmt.Sprintf("%.2f", got.TotalAmount) != fmt.Sprintf("%.2f", tt.want.TotalAmount) {
				t.Errorf("NewOrder().TotalAmount = %v, want %v", got.TotalAmount, tt.want.TotalAmount)
			}

			if got.CreatedAt.IsZero() {
				t.Error("NewOrder().CreatedAt.IsZero() == true")
			}
		})
	}
}
