package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/mnsavag/anki.git/internal/lib/log"
)

type Config struct {
	Service  `envPrefix:"SERVICE_"`
	Database `envPrefix:"DATABASE_"`

	Logger log.Config
}

type Service struct {
	Host     string `env:"HOST" envDefault:"127.0.0.1"`
	HTTPPort string `env:"HTTP_PORT" envDefault:"8080"`
}

type Database struct {
	DSN string `env:"ANKI_DSN"`
}

func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
