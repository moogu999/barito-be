package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type SQLConfig struct {
	Username        string        `env:"SQL_USERNAME, required"`
	Password        string        `env:"SQL_PASSWORD, required"`
	Host            string        `env:"SQL_HOST, required"`
	Port            string        `env:"SQL_PORT, required"`
	DatabaseName    string        `env:"SQL_DATABASE_NAME, required"`
	MaxOpenCons     int           `env:"SQL_MAX_OPEN_CONS"`
	ConnMaxLifetime time.Duration `env:"SQL_CONN_MAX_LIFETIME"`
	MaxIdleCons     int           `env:"SQL_MAX_IDLE_CONS"`
	ConnMaxIdleTime time.Duration `env:"SQL_CONN_MAX_IDLE_TIME"`
}

func LoadSQLConfig() SQLConfig {
	var cfg SQLConfig
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		msg := "unable to load sql config"
		slog.Error(msg, slog.String("error", err.Error()))
		panic(fmt.Errorf(msg))
	}
	return cfg
}
