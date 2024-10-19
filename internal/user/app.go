package user

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moogu999/barito-be/internal/infra/database/mysql"
	"github.com/moogu999/barito-be/internal/user/port"
	"github.com/moogu999/barito-be/internal/user/usecase"
)

type Dependency struct {
	DB     *sql.DB
	Router chi.Router
}

type App struct {
	HTTP http.Handler
}

func New(dep Dependency) *App {
	repo := mysql.NewUserRepository(dep.DB)
	service := usecase.NewService(repo)
	httpHandler := port.NewHTTP(dep.Router, service)
	return &App{
		HTTP: httpHandler,
	}
}
