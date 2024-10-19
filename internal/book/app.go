package book

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Dependency struct {
	DB     *sql.DB
	Router chi.Router
}

type App struct {
	Handler http.Handler
}

func NewApp(dep Dependency) *App {
	return &App{}
}
