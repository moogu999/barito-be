package config

import (
	"time"
)

type SQLConfig struct {
	Username        string        `env:"USERNAME, required"`
	Password        string        `env:"PASSWORD, required"`
	Host            string        `env:"HOST, required"`
	Port            string        `env:"PORT, required"`
	DatabaseName    string        `env:"DATABASE_NAME, required"`
	MaxOpenCons     int           `env:"MAX_OPEN_CONS, default=50"`
	ConnMaxLifetime time.Duration `env:"CONN_MAX_LIFETIME, default=1h"`
	MaxIdleCons     int           `env:"MAX_IDLE_CONS, default=10"`
	ConnMaxIdleTime time.Duration `env:"CONN_MAX_IDLE_TIME, default=15m"`
}
