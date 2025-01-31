package log

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerSuite struct {
	suite.Suite
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

func (s *LoggerSuite) TestWithField() {
	// arrange
	var buf bytes.Buffer
	conf := Config{
		LogLevel:     LevelDebugString,
		EncodingType: encodingTypeJSON,
	}
	log := new(Logger)
	log.logger = slog.New(newHandler(&conf, &buf))

	// act
	key := "service"
	value := "anki"
	log.WithField(key, value).Info("message")

	// assert
	s.Contains(buf.String(), key)
	s.Contains(buf.String(), value)
}
