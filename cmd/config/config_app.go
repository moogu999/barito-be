package config

type HTTPConfig struct {
	Port string `env:"PORT, default=8080"`
}
