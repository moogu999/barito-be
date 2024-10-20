package entity

// @TODO update find books endpoint
type Book struct {
	ID     int64
	Title  string
	Author string
	ISBN   string
	Price  float64
}
