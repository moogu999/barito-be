package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moogu999/barito-be/internal/infra/database"
	"github.com/moogu999/barito-be/internal/user"
)

func main() {
	db := database.NewSQL()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60000 * time.Millisecond))

	r.Route("/", func(r chi.Router) {
		user.New(user.Dependency{
			DB:     db,
			Router: r,
		})
	})

	// @TODO create config file
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), r)
}
