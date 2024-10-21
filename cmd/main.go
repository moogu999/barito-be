package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moogu999/barito-be/cmd/config"
	"github.com/moogu999/barito-be/internal/book"
	"github.com/moogu999/barito-be/internal/infra/database"
	"github.com/moogu999/barito-be/internal/order"
	"github.com/moogu999/barito-be/internal/user"
)

func main() {
	ctx := context.Background()

	cfg := config.Get(ctx)

	db := database.NewSQL(cfg.SQLConfig)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60000 * time.Millisecond))

	r.Route("/", func(r chi.Router) {
		user.NewApp(user.Dependency{
			DB:     db,
			Router: r,
		})
		book.NewApp(book.Dependency{
			DB:     db,
			Router: r,
		})
		order.NewApp(order.Dependency{
			DB:     db,
			Router: r,
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.HTTPConfig.Port), r)
}
