package entity

// Book represents a book in the system.
type Book struct {
	ID     int64
	Title  string
	Author string
	ISBN   string
	Price  float64
}
