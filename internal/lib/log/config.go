package log

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"local" validate:"required,oneof=local dev prod"`
}

func (conf *Config) Validate() error {
	return validator.New().Struct(conf)
}
