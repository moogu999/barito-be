package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

func NewSQL() *sql.DB {
	cfg := LoadSQLConfig()

	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&parseTime=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	))
	if err != nil {
		slog.Error("error when openning sql connection", slog.String("error", err.Error()))
		return nil
	}

	db.SetMaxOpenConns(cfg.MaxOpenCons)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleCons)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	return db
}
