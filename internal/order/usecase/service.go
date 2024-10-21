package usecase

import (
	"context"

	bookRepo "github.com/moogu999/barito-be/internal/book/domain/repository"
	"github.com/moogu999/barito-be/internal/order/domain/entity"
	"github.com/moogu999/barito-be/internal/order/domain/repository"
	userRepo "github.com/moogu999/barito-be/internal/user/domain/repository"
)

type OrderUseCase interface {
	// CreateOrder create a new order for a user.
	// If user is not found, it will return an error.
	// If book is nout found, it will return an error.
	// If purchase quantity per item is less than it will return an error.
	CreateOrder(ctx context.Context, userID int64, items []CartItem) (int64, error)

	// FindOrders return all the orders a user had made.
	FindOrders(ctx context.Context, userID int64) ([]*entity.Order, error)
}

type Service struct {
	orderRepo repository.OrderRepository
	userRepo  userRepo.UserRepository
	bookRepo  bookRepo.BookRepository
}

func NewService(orderRepo repository.OrderRepository,
	userRepo userRepo.UserRepository,
	bookRepo bookRepo.BookRepository,
) OrderUseCase {
	return &Service{
		orderRepo: orderRepo,
		userRepo:  userRepo,
		bookRepo:  bookRepo,
	}
}
