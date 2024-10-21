package entity

type OrderItem struct {
	ID     int64
	BookID int64
	Title  string
	Author string
	Qty    int
	Price  float64
}
