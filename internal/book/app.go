package book

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/book/port"
	"github.com/moogu999/barito-be/internal/book/usecase"
	"github.com/moogu999/barito-be/internal/infra/database/mysql"
)

type Dependency struct {
	DB     *sql.DB
	Router chi.Router
}

type App struct {
	Handler http.Handler
}

func NewApp(dep Dependency) *App {
	repo := mysql.NewBookRepository(dep.DB)
	service := usecase.NewService(repo)
	handler := port.NewHandler(dep.Router, service)
	return &App{
		Handler: handler,
	}
}
