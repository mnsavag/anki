package log

import (
	"io"
	"log/slog"
)

func newHandler(conf *Config, w io.Writer) slog.Handler {
	if conf.EncodingType == encodingTypeJSON {
		return slog.NewJSONHandler(w, &slog.HandlerOptions{Level: envLevelsMapping[conf.LogLevel]})
	}
	return slog.NewTextHandler(w, &slog.HandlerOptions{Level: envLevelsMapping[conf.LogLevel]})
}
