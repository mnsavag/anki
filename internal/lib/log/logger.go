package log

import (
	"fmt"
	"io"
	"os"

	"github.com/mnsavag/anki.git/internal/lib/log/fields"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
	entry  *logrus.Entry
	conf   *Config
}

func NewLogger(conf *Config) (*Logger, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	log := new(Logger)
	log.conf = conf

	log.logger = logrus.New()
	log.logger.Out = os.Stdout
	log.logger.Formatter = newFormatter(conf)
	log.logger.Level = envLevelsMapping[log.conf.LogLevel]
	log.entry = logrus.NewEntry(log.logger)

	if log.conf.Prefix != "" {
		log = log.WithField("service", log.conf.Prefix)
	}

	return log, nil
}

func (l *Logger) Info(v ...any) {
	l.entry.Info(fmt.Sprint(v...))
}

func (l *Logger) WithField(key string, value any) *Logger {
	return &Logger{
		logger: l.logger,
		entry:  l.logger.WithField(key, value),
		conf:   l.conf,
	}
}

func (l *Logger) WithFields(fields map[string]any) *Logger {
	return &Logger{
		logger: l.logger,
		entry:  l.logger.WithFields(fields),
		conf:   l.conf,
	}
}

func (l *Logger) WithError(err error) *Logger {
	if err == nil {
		return l
	}

	return l.WithField(fields.Error, err.Error())
}

func (l *Logger) SetOutput(w io.Writer) {
	l.logger.SetOutput(w)
}
