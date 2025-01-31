package log

import (
	"github.com/go-playground/validator/v10"
)

type Config struct {
	LogLevel     string `env:"LOG_LEVEL" envDefault:"INFO" validate:"required,oneof=DEBUG INFO WARN ERROR"`
	EncodingType string `env:"LOG_ENCODING_TYPE" envDefault:"JSON" validate:"required,oneof=JSON PLAIN"`
	Prefix       string `env:"LOG_PREFIX"`
}

func (conf *Config) Validate() error {
	return validator.New().Struct(conf)
}
