package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
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
	r.Use(httplog.RequestLogger(setupHTTPLogger()))
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

func setupHTTPLogger() *httplog.Logger {
	return httplog.NewLogger("barito-be", httplog.Options{
		JSON:           true,
		LogLevel:       slog.LevelInfo,
		Concise:        true,
		RequestHeaders: true,
	})
}
