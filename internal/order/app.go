package order

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/infra/database/mysql"
	"github.com/moogu999/barito-be/internal/order/port"
	"github.com/moogu999/barito-be/internal/order/usecase"
)

type Dependency struct {
	DB     *sql.DB
	Router chi.Router
}

type App struct {
	Handler http.Handler
}

func NewApp(dep Dependency) *App {
	orderRepo := mysql.NewOrderRepository(dep.DB)
	userRepo := mysql.NewUserRepository(dep.DB)
	bookRepo := mysql.NewBookRepository(dep.DB)
	service := usecase.NewService(orderRepo, userRepo, bookRepo)
	handler := port.NewHandler(dep.Router, service)
	return &App{
		Handler: handler,
	}
}
