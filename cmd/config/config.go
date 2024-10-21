package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	HTTPConfig HTTPConfig `env:", prefix=HTTP_"`
	SQLConfig  SQLConfig  `env:", prefix=SQL_"`
}

var (
	once sync.Once
	cfg  *Config
)

func Get(ctx context.Context) *Config {
	once.Do(func() {
		cfg = new(Config)
		if err := envconfig.Process(ctx, cfg); err != nil {
			panic(fmt.Errorf("unable to load config"))
		}
	})
	return cfg
}
