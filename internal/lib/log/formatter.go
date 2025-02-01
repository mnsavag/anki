package log

import (
	"github.com/sirupsen/logrus"
)

// TODO: outputer interface for PlainFormatter and JSONFormatter
func newFormatter(conf *Config) logrus.Formatter {
	if conf.EncodingType == encodingTypeJSON {
		return new(logrus.JSONFormatter)
	}
	return new(logrus.TextFormatter)
}
