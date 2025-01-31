package log

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mnsavag/anki.git/internal/lib/log/fields"
)

type Logger struct {
	logger *slog.Logger
	conf   *Config
}

func NewLogger(conf *Config) (*Logger, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	log := new(Logger)
	log.conf = conf
	log.logger = slog.New(newHandler(log.conf, os.Stdout))
	log.setPrefix()

	return log, nil
}

func (l *Logger) setPrefix() {
	if l.conf.Prefix != "" {
		l = l.WithField("service", l.conf.Prefix)
	}
}

func (l *Logger) Info(v ...any) {
	l.logger.Info(fmt.Sprint(v...))
}

func (l *Logger) WithField(key string, value any) *Logger {
	return &Logger{
		logger: l.logger.With(slog.Any(key, value)),
		conf:   l.conf,
	}
}

func (l *Logger) WithFields(fields map[string]any) *Logger {
	for key, val := range fields {
		l = l.WithField(key, val)
	}
	return l
}

func (l *Logger) WithError(err error) *Logger {
	if err == nil {
		return l
	}

	return l.WithField(fields.Error, err.Error())
}
