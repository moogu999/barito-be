package usecase

import (
	"context"
	"sync"

	bookEntity "github.com/moogu999/barito-be/internal/book/domain/entity"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	userEntity "github.com/moogu999/barito-be/internal/user/domain/entity"
)

// CartItem represents an item a user wishes to buy
type CartItem struct {
	BookID int64
	Qty    int
}

func (s *Service) CreateOrder(ctx context.Context, userID int64, cartItems []CartItem) (int64, error) {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error)
		bookIDs = getUniqueBookIDs(cartItems)
		user    *userEntity.User
		books   []*bookEntity.Book
	)

	// Fetch user and books concurrently.
	go func() {
		defer func() {
			wg.Wait()
			close(errChan)
		}()

		wg.Add(1)
		go func() {
			var userErr error
			user, userErr = s.userRepo.GetUserByID(ctx, userID)
			errChan <- userErr
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			var booksErr error
			books, booksErr = s.bookRepo.GetBooksByIDs(ctx, bookIDs)
			errChan <- booksErr
			wg.Done()
		}()
	}()

	for err := range errChan {
		if err != nil {
			return 0, err
		}
	}

	if user == nil {
		return 0, userEntity.ErrUserNotFound
	}

	// 1 or all of the books are not found.
	if len(bookIDs) != len(books) {
		return 0, bookEntity.ErrBooksNotFound
	}

	tx, err := s.orderRepo.BeginTx(ctx)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			s.orderRepo.RollbackTx(tx)
		}
	}()

	order, err := entity.NewOrder(userID, constructOrderItems(cartItems, books))
	if err != nil {
		return 0, err
	}

	err = s.orderRepo.CreateOrder(ctx, tx, &order)
	if err != nil {
		return 0, err
	}

	for _, item := range order.Items {
		err := s.orderRepo.CreateOrderItem(ctx, tx, order.ID, &item)
		if err != nil {
			return 0, err
		}
	}

	err = s.orderRepo.CommitTx(tx)
	if err != nil {
		return 0, err
	}

	return order.ID, nil
}

// getUniqueBookIDs returns all the unique book ids a user wishes to purchase.
func getUniqueBookIDs(items []CartItem) []int64 {
	idsMap := getBookIDsMap(items)

	ids := make([]int64, 0)
	for key := range idsMap {
		ids = append(ids, key)
	}

	return ids
}

// constructOrderItems combines the cart items and books
// to generate the order detail.
func constructOrderItems(cartItems []CartItem, books []*bookEntity.Book) []entity.OrderItem {
	idsMap := getBookIDsMap(cartItems)

	items := make([]entity.OrderItem, 0)
	for _, book := range books {
		items = append(items, entity.OrderItem{
			BookID: book.ID,
			Qty:    idsMap[book.ID],
			Price:  book.Price,
		})
	}

	return items
}

// getBookIDsMap returns a map with book ID as it keys
// and the purchase quantity as it values.
// It groups together item with the same book ID.
func getBookIDsMap(items []CartItem) map[int64]int {
	idsMap := make(map[int64]int, len(items))
	for _, val := range items {
		if _, ok := idsMap[val.BookID]; !ok {
			idsMap[val.BookID] = val.Qty
		} else {
			idsMap[val.BookID] += val.Qty
		}
	}

	return idsMap
}
