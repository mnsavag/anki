package log

import "log/slog"

const (
	encodingTypePlain string = "PLAIN"
	encodingTypeJSON  string = "JSON"
)

const (
	LevelDebugString = "DEBUG"
	LevelInfoString  = "INFO"
	LevelWarnString  = "WARN"
	LevelErrorString = "ERROR"
)

// DEBUG INFO WARN ERROR
var envLevelsMapping = map[string]slog.Level{
	LevelDebugString: slog.LevelDebug,
	LevelInfoString:  slog.LevelInfo,
	LevelWarnString:  slog.LevelWarn,
	LevelErrorString: slog.LevelError,
}
