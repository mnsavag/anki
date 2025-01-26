package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
	conf   *Config
}

func NewLogger(conf *Config) (*Logger, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	log := &Logger{
		logger: &slog.Logger{},
	}

	switch conf.LogLevel {
	case string(local):
		log.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case string(dev):
		log.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case string(prod):
		log.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log, nil
}
