package order

import (
	"database/sql"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNewApp(t *testing.T) {
	t.Parallel()

	r := chi.NewRouter()
	dep := Dependency{
		DB:     &sql.DB{},
		Router: r,
	}

	app := NewApp(dep)
	if app == nil {
		t.Error("app is nil")
	}
}
